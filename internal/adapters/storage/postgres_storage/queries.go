package postgres_storage

// TODO: add a 'SET ROLE' before each query to make the code more reliable and testable

var (
	// Layers
	createLayer      = `call public.create_layer(:layer_name);`
	selectLayerNames = `select substring(schema_name, 1, position('_layer' in schema_name) - 1) 
						from information_schema.schemata where schema_name like '%_layer';`
	// Users
	insertUser = `select * from public.insert_user(:username, :password, :role);`
	// selectUserPassword = `select password from public.users where name = $1 limit 1;`
	selectRole = `select users.role from public.users where users.name = (select session_user);`
	// Models
	selectAllModels = `select * from public.select_all_models(:layer_name);`
	insertModels    = `select * from public.insert_models(:layer_name, :models_array);`
	// Model elements
	selectAllModelElements = `select * from public.select_all_model_elements(:layer_name);`
	insertModelElements    = `select * from public.insert_model_elements(:layer_name, :elements_array);`
	// Properties
	selectAllProperties        = `select * from public.properties;`
	selectPropertiesByUnit     = `select * from public.select_properties_by_unit(:layer_name, :lang, :unit_text);`
	selectPropertiesIdByUnitId = `select * from public.select_properties_id_by_unit_id(:layer_name, :lang, :unit_id);`
	insertProperties           = `select * from public.insert_properties(:properties_array);`
	insertUnitProperties       = `call public.insert_unit_properties(:layer_name, :lang, :unit_id, :properties_id);`
	updateUnitProperties       = `call public.update_unit_properties(:layer_name, :lang, :unit_name, :properties_id);`
	// Contexts
	selectContextsById     = `select id, text from public.contexts where id = any($1);`
	selectContextsIdByUnit = `select * from public.select_contexts_id_by_unit(:layer_name, :lang, :unit_id);`
	insertContextUnits     = `call public.insert_context_units(:layer_name, :lang, :context_id, :units_id);`
	insertContext          = `insert into public.contexts(id, registration_date, text) overriding user value -- or overriding system value
							  values (null, now()::timestamp, $1)
							  on conflict(text) do update set text=public.contexts.text
							  returning public.contexts.id;`
	// Units
	selectAllLinkedUnits               = `select * from public.select_all_linked_units(:layer_name);`
	selectLinkedUnitsByModelsId        = `select * from public.select_linked_units_by_models_id(:layer_name, :models_id_array);`
	selectLinkedUnitsByPropertiesId    = `select * from public.select_linked_units_by_properties_id(:layer_name, :properties_id_array);`
	selectUnlinkedUnits                = `select * from public.select_unlinked_units_by_lang(:layer_name, :lang);`
	selectUnlinkedUnitsAndModelsId     = `select * from public.select_unlinked_units_by_lang_and_models_id(:layer_name, :lang, :models_id_array);`
	selectUnlinkedUnitsAndPropertiesId = `select * from public.select_unlinked_units_by_lang_and_properties_id(:layer_name, :lang, :properties_id_array);`
	insertUnits                        = `select * from public.insert_units(:layer_name, :lang, :models_id, :unit_texts);`
	linkUnits                          = `call public.link_units(:layer_name, :unit_ru, :unit_en);`
	updateUnitNames                    = `call public.update_unit_names(:layer_name, :lang, :old_name, :new_name);`
)
