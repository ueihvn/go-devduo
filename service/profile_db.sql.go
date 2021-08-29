package service

var getProfile = `
select * from profiles
`

var filterProfileByField = `
select * from profiles 
where not exists (
	select * from fields
	where fields.id in (?)
	and not exists (
		select * from profile_fields
		where profile_fields.field_id = fields.id
		and profile_fields.profile_user_id = profiles.user_id
	)
)
`
var filterProfileByTech = `
select * from profiles
where not exists (
	select * from technologies
	where technologies.id in (?)
	and not exists (
		select * from profile_technologies
		where profile_technologies.technology_id = technologies.id
		and profile_technologies.profile_user_id = profiles.user_id
	)
)
`

var filterProfileByFieldTech = `
select * from profiles 
where not exists (
	select * from fields
	where fields.id in (?)
	and not exists (
		select * from profile_fields
		where profile_fields.field_id = fields.id
		and profile_fields.profile_user_id = profiles.user_id
	)
) and profiles.user_id in 
(select profiles.user_id from profiles
	where not exists (
		select * from technologies
		where technologies.id in (?)
		and not exists (
			select * from profile_technologies
			where profile_technologies.technology_id = technologies.id
			and profile_technologies.profile_user_id = profiles.user_id
		)
	)
)
`
