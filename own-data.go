package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"sort"
)

func PrepareExtraData() {
	prepareCubers()
	prepareRankingEntries()
}

// ----------------------- (helpers) -----------------------------

func singles(result WcaResult) []int32 {
	out := []int32{}
	for _, value := range []int32{result.Value1, result.Value2, result.Value3, result.Value4, result.Value5} {
		if value != 0 {
			out = append(out, value)
		}
	}
	return out
}

func gzipIt(data []byte) []byte {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	writer.Write(data)
	writer.Close()
	return buffer.Bytes()
}

type NameAndId struct {
	Name str32 `json:"name"`
	Id   str32 `json:"id"`
}

// ----------------------- Cubers -----------------------------

type Cuber struct {
	Id               str32 `json:"id"`
	Name             str32 `json:"name"`
	CountryId        str32 `json:"countryId"`
	Gender           str32 `json:"gender"`
	CompetitionCount int32 `json:"competitionCount"`
	competitions     map[str32]bool
	results          []CuberResult
	resultsJsonGzip  []byte
}

type CuberResult struct {
	Competition   NameAndId `json:"competition"`
	EventId       str32     `json:"eventId"`
	RoundId       str32     `json:"roundId"`
	Pos           int32     `json:"pos"`
	Best          int32     `json:"best"`
	Average       int32     `json:"average,omitempty"`
	Singles       []int32   `json:"singles"`
	BestMarker    str32     `json:"bestMarker,omitempty"`
	AverageMarker str32     `json:"averageMarker,omitempty"`
}

func newCuberResult(result WcaResult) CuberResult {
	return CuberResult{
		Competition:   NameAndId{result.CompetitionId, result.CompetitionId},
		EventId:       result.EventId,
		RoundId:       result.RoundId,
		Pos:           result.Pos,
		Best:          result.Best,
		Average:       result.Average,
		Singles:       singles(result),
		BestMarker:    result.RegionalSingleRecord,
		AverageMarker: result.RegionalAverageRecord,
	}
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

	// Collect each cubers competitions and results
	for _, result := range wcaResults {
		cuber := cubers[result.PersonId]
		cuber.competitions[result.CompetitionId] = true
		cuber.results = append(cuber.results, newCuberResult(result))
		cubers[result.PersonId] = cuber
	}

	// Count each cuber's competitions and prepare Sebastien's results answer
	for i := range cubers {
		cuber := cubers[i]
		cuber.CompetitionCount = int32(len(cuber.competitions))
		if cuber.Id.String() == "2008AURO01" {
			j, _ := json.Marshal(cuber.results)
			cuber.resultsJsonGzip = gzipIt(j)
		}
		cubers[i] = cuber
	}
}

// ----------------------- RankingEntries -----------------------------

type RankingEntry struct {
	Cuber       NameAndId `json:"cuber"`
	Value       int32     `json:"value"`
	Country     str32     `json:"country"`
	Competition NameAndId `json:"competition"`
}

type RankingEntryKey struct {
	EventId   str32
	CuberId   str32
	CountryId str32
}

var rankingEntries map[str32][]RankingEntry

type ByValue []WcaResult

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Best < a[j].Best }

func prepareRankingEntries() {

	// Sort all results by Best (TODO: and chronologically)
	sort.Sort(ByValue(wcaResults))

	// Collect the ranking entries, grouped by event
	rankingEntries = make(map[str32][]RankingEntry)
	seen := map[RankingEntryKey]bool{}
	for _, result := range wcaResults {
		if result.Best > 0 {
			key := RankingEntryKey{result.EventId, result.PersonId, result.PersonCountryId}
			if !seen[key] {
				seen[key] = true
				entry := RankingEntry{
					Cuber:       NameAndId{result.PersonName, result.PersonId},
					Value:       result.Best,
					Country:     result.PersonCountryId,
					Competition: NameAndId{result.CompetitionId, result.CompetitionId},
				}
				rankingEntries[result.EventId] = append(rankingEntries[result.EventId], entry)
			}
		}
	}
}
