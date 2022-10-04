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

create or replace
function public.get_next(id int)
	returns integer
	language plpgsql
as $function$
	begin	
		return (select count(*) from Bookmark where "owner" = $1) + 1;
end;
$function$;

create or replace
function concat(id int)
	returns integer
	language plpgsql
as $function$
	begin
		return (select cast(cast(get_next(id) as text) || '0'::text || cast(random_between(1000, 9000) as text) as numeric(24, 0)));
end;
$function$;