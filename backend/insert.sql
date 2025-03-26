with new_inv_items as (
    select
        generate_series(1, 1) as item_count,
        skin.id as skin_id,
        skin.wear_min,
        skin.wear_max,
        skin.can_be_stattrak
    from skins skin
    where skin.Rarity = 'Consumer' 
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
    nextval('inventory_id_seq') as id,
    (select id from users where username='electronicbob'),
    skin_id,
    (generate_skin_wear(wear_min, wear_max)).wear_str,
    (generate_skin_wear(wear_min, wear_max)).wear_num,
    (random_price()) as price,
    can_be_stattrak and (random() < 0.1) as is_stattrak,
    now() as created_at
from new_inv_items;
