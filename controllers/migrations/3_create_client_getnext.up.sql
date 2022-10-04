CREATE OR REPLACE FUNCTION get_next()
	RETURNS int
	LANGUAGE plpgsql
AS $function$
	BEGIN
		RETURN (SELECT count(*) FROM Bookmark) + 1;
	END;
$function$;