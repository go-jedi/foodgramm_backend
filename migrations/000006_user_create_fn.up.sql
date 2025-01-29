CREATE OR REPLACE FUNCTION public.user_create(_src JSON) RETURNS users
    SECURITY DEFINER
    LANGUAGE plpgsql
AS
$$
DECLARE
    _u users;
BEGIN
    INSERT INTO users(
        telegram_id,
        username,
        first_name,
        last_name
    ) VALUES(
        _src->>'telegram_id',
        _src->>'username',
        _src->>'first_name',
        _src->>'last_name'
    ) RETURNING * INTO _u;

    INSERT INTO user_excluded_products_table(
        user_id,
        telegram_id
    ) VALUES(
        _u.id,
        _u.telegram_id
    );

    RETURN _u;
END;
$$;