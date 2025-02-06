CREATE OR REPLACE FUNCTION public.recipe_create(_at INTEGER, _src JSON) RETURNS recipes
    SECURITY DEFINER
    LANGUAGE plpgsql
AS
$$
DECLARE
    _r recipes;
BEGIN
    INSERT INTO recipes(
        telegram_id,
        title,
        content
    ) VALUES(
        _src->>'telegram_id',
        _src->>'title',
        (_src->>'content')::JSONB
    ) RETURNING * INTO _r;

    IF _r.id ISNULL THEN
        RAISE EXCEPTION 'there was an error creating the recipe';
    END IF;

    IF _at = 2 THEN
        UPDATE user_free_recipes SET
            free_recipes_used = free_recipes_used - 1,
            updated_at = NOW()
        WHERE telegram_id = _src->>'telegram_id';
    END IF;

    RETURN _r;
END;
$$;