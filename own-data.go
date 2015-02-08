package main

func PrepareExtraData() {
	prepareCubers()
}

// ----------------------- Cubers -----------------------------

type Cuber struct {
	Id               str32 `json:"id"`
	Name             str32 `json:"name"`
	CountryId        str32 `json:"countryId"`
	Gender           str32 `json:"gender"`
	CompetitionCount int32 `json:"competitionCount"`
	competitions     map[str32]bool
}

var cubers = map[str32]Cuber{}

func getCuber(id string) (cuber Cuber, ok bool) {
	cuber, ok = cubers[getStr32(id)]
	return
}

func prepareCubers() {

	// Convert persons with subid==1 to cubers.
	for _, person := range wcaPersons {
		if person.Subid == 1 {
			cubers[person.Id] = Cuber{
				Id:           person.Id,
				Name:         person.Name,
				CountryId:    person.CountryId,
				Gender:       person.Gender,
				competitions: map[str32]bool{},
			}
		}
	}

	// Collect and count each cubers competitions
	for _, result := range wcaResults {
		cubers[result.PersonId].competitions[result.CompetitionId] = true
	}
	for i := range cubers {
		cuber := cubers[i]
		cuber.CompetitionCount = int32(len(cuber.competitions))
		cubers[i] = cuber
	}
}
