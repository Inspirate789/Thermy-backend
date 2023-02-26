select *
from public.properties
where id in (select property_id
             from $1_layer.properties_and_units_$2
             where unit_id = (select id
                              from $1_layer.units_$2
                              where text = $3
                              limit 1 -- because text has unique constraint
                             )
            );
