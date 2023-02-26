select $1_layer.units_ru.id as unit_ru_id,
       $1_layer.units_ru.model_id as unit_ru_model_id,
       $1_layer.units_ru.registration_date as unit_ru_registration_date,
       $1_layer.units_ru.text as unit_ru_text,
       $1_layer.units_en.id as unit_en_id,
       $1_layer.units_en.model_id as unit_en_model_id,
       $1_layer.units_en.registration_date as unit_en_registration_date,
       $1_layer.units_en.text as unit_en_text
from  $1_layer.units_ru
    inner join $1_layer.units_ru_and_en on $1_layer.units_ru.id = unit_ru_id
    inner join $1_layer.units_en on $1_layer.units_en.id = unit_en_id;
