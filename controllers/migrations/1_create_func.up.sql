CREATE OR REPLACE FUNCTION random_between(low integer, high integer)
	RETURNS integer
	LANGUAGE plpgsql
	STRICT
AS $function$
declare
	output int := 0;
begin
	output := floor(random() * (high-low + 1) + low);
	RETURN (output);
END;
$function$;