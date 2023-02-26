select *
from  $1_layer.units_$2
where id not in (select unit_$2_id
                 from $1_layer.units_ru_and_en
                );
