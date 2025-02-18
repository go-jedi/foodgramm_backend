CREATE OR REPLACE FUNCTION public.promo_code_create(_src JSON) RETURNS promo_codes
    SECURITY DEFINER
    LANGUAGE plpgsql
AS
$$
DECLARE
    _pc promo_codes;
BEGIN
    SELECT *
    FROM promo_codes
    WHERE code = _src->>'code'
    INTO _pc;

    IF _pc.id IS NOT NULL THEN
       RAISE EXCEPTION 'promo code with this code already exists';
    END IF;

    INSERT INTO promo_codes(
        code,
        discount_percent,
        max_uses_allowed,
        is_reusable,
        valid_until
    ) VALUES(
        _src->>'code',
        (_src->>'discount_percent')::INTEGER,
        (_src->>'max_uses_allowed')::INTEGER,
        (_src->>'is_reusable')::BOOLEAN,
        (_src->>'valid_until')::TIMESTAMP
    ) RETURNING * INTO _pc;

    RETURN _pc;
END;
$$;