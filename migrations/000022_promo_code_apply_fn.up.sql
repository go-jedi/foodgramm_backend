CREATE OR REPLACE FUNCTION public.promo_code_apply(_code TEXT, _telegram_id TEXT) RETURNS JSONB
    SECURITY DEFINER
    LANGUAGE plpgsql
AS
$$
DECLARE
    _pc promo_codes;
    _result JSONB;
BEGIN
    SELECT *
    FROM promo_codes
    WHERE code = _code
    AND valid_from <= NOW()
    AND (valid_until IS NULL OR valid_until >= NOW())
    AND (max_uses_allowed = -1 OR amount_used < max_uses_allowed)
    AND (is_reusable OR amount_used = 0)
    AND NOT EXISTS (
        SELECT 1
        FROM promo_code_uses
        WHERE promo_code_id = promo_codes.id
        AND telegram_id = _telegram_id
    ) INTO _pc;

    -- если промокод не доступен, то exception, а иначе применяем его
    IF _pc.id IS NULL THEN
        RAISE EXCEPTION 'the promo code is either invalid, expired, or has reached its usage limit. code: %, telegram_id: %', _code, _telegram_id;
    ELSE
        -- увеличиваем количество использований
        UPDATE promo_codes
        SET amount_used = amount_used + 1
        WHERE id = _pc.id;

        -- фиксируем использование промокода пользователем
        INSERT INTO promo_code_uses(
            promo_code_id,
            telegram_id
        ) VALUES(
            _pc.id,
            _telegram_id
        );
    END IF;

    -- формируем JSONB-объект с результатом
    _result := jsonb_build_object(
        'telegram_id', _telegram_id,
        'discount_percent', _pc.discount_percent
    );

    -- возвращаем результат
    RETURN _result;
END;
$$;