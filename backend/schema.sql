create type status as enum ('Active', 'In Progress', 'Completed');

create table if not exists skins (
    id int primary key not null,
    name text,
    rarity text,
    collection text,
    wear_min float,
    wear_max float,
    can_be_stattrak boolean,
    image_key text
);

create table if not exists users (
    id uuid primary key not null,
    username text,
    email text,
    hash text,
    created_at timestamp,
    refresh_token_version int default 1
);

create table if not exists inventory (
    id int primary key not null,
    user_id uuid not null,
    skin_id int not null,
    wear_str text,
    wear_num float,
    price float,
    is_stattrak boolean,
    was_won boolean default false,
    created_at timestamp,

    foreign key (user_id) references users(id),
    foreign key (skin_id) references skins(id)
);

create table if not exists tradeups (
    id int primary key not null,
    rarity text,
    current_status status,
    winner uuid,
    stop_time timestamp,

    foreign key (winner) references users(id)
);

create table if not exists tradeups_skins (
    tradeup_id int not null,
    inv_id int not null,

    primary key(tradeup_id, inv_id),
    foreign key (tradeup_id) references tradeups(id),
    foreign key (inv_id) references inventory(id)
);

create table if not exists crates (
    id int primary key generated always as identity not null,
    name text not null,
    rarity text not null,
    image_key text,
    cost numeric not null
);

create table if not exists crate_skins (
    crate_id int not null,
    skin_id int not null,

    primary key(crate_id, skin_id),
    foreign key(crate_id) references crates(id),
    foreign key(skin_id) references skins(id)
);
