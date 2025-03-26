--create extension if not exists hstore;

--create table wear_ranges (
--    ranges hstore
--);

--insert into wear_ranges(ranges) values (
--    '"Factory New" => "0.00-0.07", "Minimal Wear" => "0.07-0.15", "Field-Tested" => "0.15-0.38", "Well-Worn" => "0.38-0.45", "Battle-Scarred" => "0.45-1.00"'
--)

create or replace function random_float_in_range(range_str text) returns numeric as $$
declare
    min_val numeric;
    max_val numeric;
begin
    min_val := split_part(range_str, '-', 1)::numeric;
    max_val := split_part(range_str, '-', 2)::numeric;

    return min_val + (random() * (max_val - min_val));
end;
$$ language plpgsql;

create or replace function random_price() returns float as $$
begin
    return 0 + (random() * (100));
end;
$$ language plpgsql;

create or replace function generate_skin_wear(min_wear float, max_wear float)
returns table (wear_num float, wear_str text) as $$
declare
    random_wear float;
begin
    random_wear := min_wear + (random() * (max_wear - min_wear));

    RETURN QUERY
    SELECT 
        random_wear AS wear_num,
        CASE 
            WHEN random_wear BETWEEN 0.00 AND 0.07 THEN 'Factory New'
            WHEN random_wear BETWEEN 0.07 AND 0.15 THEN 'Minimal Wear'
            WHEN random_wear BETWEEN 0.15 AND 0.38 THEN 'Field-Tested'
            WHEN random_wear BETWEEN 0.38 AND 0.45 THEN 'Well-Worn'
            WHEN random_wear BETWEEN 0.45 AND 1.00 THEN 'Battle-Scarred'
            ELSE 'Unknown'
        END AS wear_str;
end;
$$ language plpgsql;

create or replace function open_crate(
    p_user_id uuid,
    p_rarity text,
    p_quantity int default 5
) returns setof inventory as $$
declare
    result_record inventory%rowtype;
begin
    for result_record in
        with random_skins as (
            select
                id as skin_id,
                wear_min,
                wear_max,
                can_be_stattrak
            from
                skins
            where
                rarity = p_rarity
            order by
                random()
            limit
                p_quantity
        )
        insert into inventory (
            id,
            user_id,
            skin_id,
            wear_str,
            wear_num,
            price,
            is_stattrak,
            created_at
        )
        select
            nextval('inventory_id_seq'),
            p_user_id,
            skin_id,
            (generate_skin_wear(wear_min, wear_max)).wear_str,
            (generate_skin_wear(wear_min, wear_max)).wear_num,
            (random_price()),
            can_be_stattrak and (random() < 0.1),
            now()
        from random_skins
        returning *
    loop
        return next result_record;
    end loop;

    return;
end;
$$ language plpgsql;
