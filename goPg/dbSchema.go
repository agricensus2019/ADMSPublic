package goPg

import (
	"fmt"
	"time"

	//	ag "AgriTracking/model"

	"AgriInject/jsonSchema"
)

// Root
type Root struct {
	CollectionName string
	Questionnaire  []*Questionnaire
	TallySheet     *TallySheet
}

func FromJsonData(json jsonSchema.Root) (root Root, err error) {
	var questionnaires []*Questionnaire
	for _, q := range json.Questionnaire {
		if len(json.TallySheet.TallySheetNo) != 6 {
			return root, fmt.Errorf("TallySheet.TallySheetNo hasn't the good lenght (6)")
		}
		questionnaires = append(questionnaires, &Questionnaire{
			QuestionnaireNum:   q.QuestionnaireNum,
			QuestionnaireEmpty: q.QuestionnaireEmpty,
			FormName:           q.FormName,
			BookletNumber:      json.TallySheet.TallySheetNo[0:3] + "." + json.TallySheet.TallySheetNo[3:6],
			TallySheetNo:       json.TallySheet.TallySheetNo,
			AgriLaborCode:      q.HouseHoldInformation.AgriLaborCode,
			FatherName:         q.HouseHoldInformation.FatherName,
			OwnerGender:        q.HouseHoldInformation.OwnerGender,
			OwnerName:          q.HouseHoldInformation.OwnerName,
			HouseSerial:        q.HouseHoldInformation.HouseSerial,
			IsFishingRelated:   q.HouseHoldInformation.IsFishingRelated,
			MobileNo:           q.HouseHoldInformation.MobileNo,
			OwnerEduCode:       q.HouseHoldInformation.OwnerEduCode,
			OwnerProfesCode:    q.HouseHoldInformation.OwnerProfesCode,
			IsUnion:            q.Location.IsUnion,
			GeocodeID: fmt.Sprintf("%02d.%02d.%03d.%03d.%03d.%d",
				q.Location.District,
				q.Location.Upazila,
				q.Location.Union,
				q.Location.Mouza,
				q.Location.CountingArea,
				q.Location.RMO),
			Village:                           q.Location.Village,
			BusinessEntityLand:                q.Table10.BusinessEntityLand.AsFloat(),
			CultivatedLand:                    q.Table10.CultivatedLand.AsFloat(),
			IrrigationLand:                    q.Table10.IrrigationLand.AsFloat(),
			LandGiven:                         q.Table10.LandGiven.AsFloat(),
			LandTaken:                         q.Table10.LandTaken.AsFloat(),
			NurseryLand:                       q.Table10.NurseryLand.AsFloat(),
			OperatingLand:                     q.Table10.OperatingLand.AsFloat(),
			PermanentCultivatedLand:           q.Table10.PermanentCultivatedLand.AsFloat(),
			PermanentUnusedLand:               q.Table10.PermanentUnusedLand.AsFloat(),
			PondLandNotBlank:                  q.Table10.PondLandNotBlank.AsFloat(),
			RecentUnusedLand:                  q.Table10.RecentUnusedLand.AsFloat(),
			ResidenceLand:                     q.Table10.ResidenceLand.AsFloat(),
			SaltCultivationLand:               q.Table10.SaltCultivationLand.AsFloat(),
			SinkChannelLand:                   q.Table10.SinkChannelLand.AsFloat(),
			TempCultivatedLand:                q.Table10.TempCultivatedLand.AsFloat(),
			TotalLand:                         q.Table10.TotalLand.AsFloat(),
			UncultivatedLand:                  q.Table10.UncultivatedLand.AsFloat(),
			CreekLand:                         q.Table11.CreekLand.AsFloat(),
			FishCageCultiveLand:               q.Table11.FishCageCultiveLand.AsFloat(),
			FishCultivationLand:               q.Table11.FishCultivationLand.AsFloat(),
			FishSaltCultiveLand:               q.Table11.FishSaltCultiveLand.AsFloat(),
			MixedCultivationLand:              q.Table11.MixedCultivationLand.AsFloat(),
			PaddyCultivationLand:              q.Table11.PaddyCultivationLand.AsFloat(),
			PondLand:                          q.Table11.PondLand.AsFloat(),
			CockAtFarm:                        int(q.Table12.Cock.AtFarm),
			CockAtHome:                        int(q.Table12.Cock.AtHome),
			DuckAtFarm:                        int(q.Table12.Duck.AtFarm),
			DuckAtHome:                        int(q.Table12.Duck.AtHome),
			PigeonAtFarm:                      int(q.Table12.Pigeon.AtFarm),
			PigeonAtHome:                      int(q.Table12.Pigeon.AtHome),
			QuailAtFarm:                       int(q.Table12.Quail.AtFarm),
			QuailAtHome:                       int(q.Table12.Quail.AtHome),
			TurkeyAtFarm:                      int(q.Table12.Turkey.AtFarm),
			TurkeyAtHome:                      int(q.Table12.Turkey.AtHome),
			BuffaloAtFarm:                     int(q.Table13.Buffalo.AtFarm),
			BuffaloAtHome:                     int(q.Table13.Buffalo.AtHome),
			CowAtFarm:                         int(q.Table13.Cow.AtFarm),
			CowAtHome:                         int(q.Table13.Cow.AtHome),
			GoatAtFarm:                        int(q.Table13.Goat.AtFarm),
			GoatAtHome:                        int(q.Table13.Goat.AtHome),
			HorseAtFarm:                       int(q.Table13.Horse.AtFarm),
			HorseAtHome:                       int(q.Table13.Horse.AtHome),
			PigAtFarm:                         int(q.Table13.Pig.AtFarm),
			PigAtHome:                         int(q.Table13.Pig.AtHome),
			SheepAtFarm:                       int(q.Table13.Sheep.AtFarm),
			SheepAtHome:                       int(q.Table13.Sheep.AtHome),
			TempAddCrop01Name:                 q.Table14.TempAddCrop01.AddCrop.CropName,
			TempAddCrop01Id:                   int(q.Table14.TempAddCrop01.AddCrop.Id),
			TempAddCrop01Surface:              q.Table14.TempAddCrop01.Count.AsFloat(),
			TempAddCrop02Name:                 q.Table14.TempAddCrop02.AddCrop.CropName,
			TempAddCrop02Id:                   int(q.Table14.TempAddCrop02.AddCrop.Id),
			TempAddCrop02Surface:              q.Table14.TempAddCrop02.Count.AsFloat(),
			TempAddCrop03Name:                 q.Table14.TempAddCrop03.AddCrop.CropName,
			TempAddCrop03Id:                   int(q.Table14.TempAddCrop03.AddCrop.Id),
			TempAddCrop03Surface:              q.Table14.TempAddCrop03.Count.AsFloat(),
			TempAddCrop04Name:                 q.Table14.TempAddCrop04.AddCrop.CropName,
			TempAddCrop04Id:                   int(q.Table14.TempAddCrop04.AddCrop.Id),
			TempAddCrop04Surface:              q.Table14.TempAddCrop04.Count.AsFloat(),
			TempAddCrop05Name:                 q.Table14.TempAddCrop05.AddCrop.CropName,
			TempAddCrop05Id:                   int(q.Table14.TempAddCrop05.AddCrop.Id),
			TempAddCrop05Surface:              q.Table14.TempAddCrop05.Count.AsFloat(),
			TempAddCrop06Name:                 q.Table14.TempAddCrop06.AddCrop.CropName,
			TempAddCrop06Id:                   int(q.Table14.TempAddCrop06.AddCrop.Id),
			TempAddCrop06Surface:              q.Table14.TempAddCrop06.Count.AsFloat(),
			TempAddCrop07Name:                 q.Table14.TempAddCrop07.AddCrop.CropName,
			TempAddCrop07Id:                   int(q.Table14.TempAddCrop07.AddCrop.Id),
			TempAddCrop07Surface:              q.Table14.TempAddCrop07.Count.AsFloat(),
			TempAddCrop08Name:                 q.Table14.TempAddCrop08.AddCrop.CropName,
			TempAddCrop08Id:                   int(q.Table14.TempAddCrop08.AddCrop.Id),
			TempAddCrop08Surface:              q.Table14.TempAddCrop08.Count.AsFloat(),
			TempAddCrop09Name:                 q.Table14.TempAddCrop09.AddCrop.CropName,
			TempAddCrop09Id:                   int(q.Table14.TempAddCrop09.AddCrop.Id),
			TempAddCrop09Surface:              q.Table14.TempAddCrop09.Count.AsFloat(),
			TempCrop101Aush:                   q.Table14.TempCrop101Aush.AsFloat(),
			TempCrop102Almond:                 q.Table14.TempCrop102Almond.AsFloat(),
			TempCrop103Boro:                   q.Table14.TempCrop103Boro.AsFloat(),
			TempCrop104Wheat:                  q.Table14.TempCrop104Wheat.AsFloat(),
			TempCrop105Maize:                  q.Table14.TempCrop105Maize.AsFloat(),
			TempCrop112Lentil:                 q.Table14.TempCrop112Lentil.AsFloat(),
			TempCrop113Safran:                 q.Table14.TempCrop113Safran.AsFloat(),
			TempCrop114Moog:                   q.Table14.TempCrop114Moog.AsFloat(),
			TempCrop121Potato:                 q.Table14.TempCrop121Potato.AsFloat(),
			TempCrop122Brinjal:                q.Table14.TempCrop122Brinjal.AsFloat(),
			TempCrop123Radish:                 q.Table14.TempCrop123Radish.AsFloat(),
			TempCrop124Bean:                   q.Table14.TempCrop124Bean.AsFloat(),
			TempCrop125Tomato:                 q.Table14.TempCrop125Tomato.AsFloat(),
			TempCrop127Multitude:              q.Table14.TempCrop127Multitude.AsFloat(),
			TempCrop128ladyfinger:             q.Table14.TempCrop128ladyfinger.AsFloat(),
			TempCrop129Cucumber:               q.Table14.TempCrop129Cucumber.AsFloat(),
			TempCrop130BitterGourd:            q.Table14.TempCrop130BitterGourd.AsFloat(),
			TempCrop131Ground:                 q.Table14.TempCrop131Ground.AsFloat(),
			TempCrop132Pumplin:                q.Table14.TempCrop132Pumplin.AsFloat(),
			TempCrop134Cauliflower:            q.Table14.TempCrop134Cauliflower.AsFloat(),
			TempCrop135Cabbage:                q.Table14.TempCrop135Cabbage.AsFloat(),
			TempCrop157Onion:                  q.Table14.TempCrop157Onion.AsFloat(),
			TempCrop158Garlic:                 q.Table14.TempCrop158Garlic.AsFloat(),
			TempCrop159Ginger:                 q.Table14.TempCrop159Ginger.AsFloat(),
			TempCrop160Turmeric:               q.Table14.TempCrop160Turmeric.AsFloat(),
			TempCrop161Chili:                  q.Table14.TempCrop161Chili.AsFloat(),
			TempCrop167Mustard:                q.Table14.TempCrop167Mustard.AsFloat(),
			TempCrop169Nuts:                   q.Table14.TempCrop169Nuts.AsFloat(),
			TempCrop175Banana:                 q.Table14.TempCrop175Banana.AsFloat(),
			TempCrop176Papaya:                 q.Table14.TempCrop176Papaya.AsFloat(),
			TempCrop177WMelon:                 q.Table14.TempCrop177WMelon.AsFloat(),
			TempCrop179PineApple:              q.Table14.TempCrop179PineApple.AsFloat(),
			TempCrop182Jute:                   q.Table14.TempCrop182Jute.AsFloat(),
			TempCrop185SugarCane:              q.Table14.TempCrop185SugarCane.AsFloat(),
			TempCrop203Seeded:                 q.Table14.TempCrop203Seeded.AsFloat(),
			FixedCrop501MongoSurface:          q.Table15.FixedCrop501Mongo.AsFloat(),
			FixedCrop501MongoCount:            int(q.Table15.FixedCrop501Mongo.TreeCount),
			FixedCrop502BerrySurface:          q.Table15.FixedCrop502Berry.AsFloat(),
			FixedCrop502BerryCount:            int(q.Table15.FixedCrop502Berry.TreeCount),
			FixedCrop503JackFruitSurface:      q.Table15.FixedCrop503JackFruit.AsFloat(),
			FixedCrop503JackFruitCount:        int(q.Table15.FixedCrop503JackFruit.TreeCount),
			FixedCrop504LitchiSurface:         q.Table15.FixedCrop504Litchi.AsFloat(),
			FixedCrop504LitchiCount:           int(q.Table15.FixedCrop504Litchi.TreeCount),
			FixedCrop505GuavaSurface:          q.Table15.FixedCrop505Guava.AsFloat(),
			FixedCrop505GuavaCount:            int(q.Table15.FixedCrop505Guava.TreeCount),
			FixedCrop506CoconutSurface:        q.Table15.FixedCrop506Coconut.AsFloat(),
			FixedCrop506CoconutCount:          int(q.Table15.FixedCrop506Coconut.TreeCount),
			FixedCrop507PlumSurface:           q.Table15.FixedCrop507Plum.AsFloat(),
			FixedCrop507PlumCount:             int(q.Table15.FixedCrop507Plum.TreeCount),
			FixedCrop508HogPlumSurface:        q.Table15.FixedCrop508HogPlum.AsFloat(),
			FixedCrop508HogPlumCount:          int(q.Table15.FixedCrop508HogPlum.TreeCount),
			FixedCrop510DateSurface:           q.Table15.FixedCrop510Date.AsFloat(),
			FixedCrop510DateCount:             int(q.Table15.FixedCrop510Date.TreeCount),
			FixedCrop511PalmyraSurface:        q.Table15.FixedCrop511Palmyra.AsFloat(),
			FixedCrop511PalmyraCount:          int(q.Table15.FixedCrop511Palmyra.TreeCount),
			FixedCrop512BellSurface:           q.Table15.FixedCrop512Bell.AsFloat(),
			FixedCrop512BellCount:             int(q.Table15.FixedCrop512Bell.TreeCount),
			FixedCrop521AverrhoaSurface:       q.Table15.FixedCrop521Averrhoa.AsFloat(),
			FixedCrop521AverrhoaCount:         int(q.Table15.FixedCrop521Averrhoa.TreeCount),
			FixedCrop522TamarindSurface:       q.Table15.FixedCrop522Tamarind.AsFloat(),
			FixedCrop522TamarindCount:         int(q.Table15.FixedCrop522Tamarind.TreeCount),
			FixedCrop523LemonSurface:          q.Table15.FixedCrop523Lemon.AsFloat(),
			FixedCrop523LemonCount:            int(q.Table15.FixedCrop523Lemon.TreeCount),
			FixedCrop524GrapeSurface:          q.Table15.FixedCrop524Grape.AsFloat(),
			FixedCrop524GrapeCount:            int(q.Table15.FixedCrop524Grape.TreeCount),
			FixedCrop538BattleLeafSurface:     q.Table15.FixedCrop538BattleLeaf.AsFloat(),
			FixedCrop538BattleLeafCount:       int(q.Table15.FixedCrop538BattleLeaf.TreeCount),
			FixedCrop539ArecaCatechuSurface:   q.Table15.FixedCrop539ArecaCatechu.AsFloat(),
			FixedCrop539ArecaCatechuCount:     int(q.Table15.FixedCrop539ArecaCatechu.TreeCount),
			FixedCrop546BambooSurface:         q.Table15.FixedCrop546Bamboo.AsFloat(),
			FixedCrop546BambooCount:           int(q.Table15.FixedCrop546Bamboo.TreeCount),
			FixedCrop548HardwoodSurface:       q.Table15.FixedCrop548Hardwood.AsFloat(),
			FixedCrop548HardwoodCount:         int(q.Table15.FixedCrop548Hardwood.TreeCount),
			FixedCrop549RainTreeSurface:       q.Table15.FixedCrop549RainTree.AsFloat(),
			FixedCrop549RainTreeCount:         int(q.Table15.FixedCrop549RainTree.TreeCount),
			FixedCrop550MahoganySurface:       q.Table15.FixedCrop550Mahogany.AsFloat(),
			FixedCrop550MahoganyCount:         int(q.Table15.FixedCrop550Mahogany.TreeCount),
			FixedCrop551TectonaGrandisSurface: q.Table15.FixedCrop551TectonaGrandis.AsFloat(),
			FixedCrop551TectonaGrandisCount:   int(q.Table15.FixedCrop551TectonaGrandis.TreeCount),
			FixedCrop572NeemSurface:           q.Table15.FixedCrop572Neem.AsFloat(),
			FixedCrop572NeemCount:             int(q.Table15.FixedCrop572Neem.TreeCount),
			FixedCropAddTree01Surface:         q.Table15.FixedCropAddTree01.Count.AsFloat(),
			FixedCropAddTree01Name:            q.Table15.FixedCropAddTree01.AddCrop.CropName,
			FixedCropAddTree01Count:           int(q.Table15.FixedCropAddTree01.Count.TreeCount),
			FixedCropAddTree01Id:              int(q.Table15.FixedCropAddTree01.AddCrop.Id),
			FixedCropAddTree02Surface:         q.Table15.FixedCropAddTree02.Count.AsFloat(),
			FixedCropAddTree02Name:            q.Table15.FixedCropAddTree02.AddCrop.CropName,
			FixedCropAddTree02Count:           int(q.Table15.FixedCropAddTree02.Count.TreeCount),
			FixedCropAddTree02Id:              int(q.Table15.FixedCropAddTree02.AddCrop.Id),
			FixedCropAddTree03Surface:         q.Table15.FixedCropAddTree03.Count.AsFloat(),
			FixedCropAddTree03Name:            q.Table15.FixedCropAddTree03.AddCrop.CropName,
			FixedCropAddTree03Count:           int(q.Table15.FixedCropAddTree03.Count.TreeCount),
			FixedCropAddTree03Id:              int(q.Table15.FixedCropAddTree03.AddCrop.Id),
			FixedCropAddTree04Surface:         q.Table15.FixedCropAddTree04.Count.AsFloat(),
			FixedCropAddTree04Name:            q.Table15.FixedCropAddTree04.AddCrop.CropName,
			FixedCropAddTree04Count:           int(q.Table15.FixedCropAddTree04.Count.TreeCount),
			FixedCropAddTree04Id:              int(q.Table15.FixedCropAddTree04.AddCrop.Id),
			FixedCropAddTree05Surface:         q.Table15.FixedCropAddTree05.Count.AsFloat(),
			FixedCropAddTree05Name:            q.Table15.FixedCropAddTree05.AddCrop.CropName,
			FixedCropAddTree05Count:           int(q.Table15.FixedCropAddTree05.Count.TreeCount),
			FixedCropAddTree05Id:              int(q.Table15.FixedCropAddTree05.AddCrop.Id),
			FixedCropAddTree06Surface:         q.Table15.FixedCropAddTree06.Count.AsFloat(),
			FixedCropAddTree06Name:            q.Table15.FixedCropAddTree06.AddCrop.CropName,
			FixedCropAddTree06Count:           int(q.Table15.FixedCropAddTree06.Count.TreeCount),
			FixedCropAddTree06Id:              int(q.Table15.FixedCropAddTree06.AddCrop.Id),
			FixedCropAddTree07Surface:         q.Table15.FixedCropAddTree07.Count.AsFloat(),
			FixedCropAddTree07Name:            q.Table15.FixedCropAddTree07.AddCrop.CropName,
			FixedCropAddTree07Count:           int(q.Table15.FixedCropAddTree07.Count.TreeCount),
			FixedCropAddTree07Id:              int(q.Table15.FixedCropAddTree07.AddCrop.Id),
			FixedCropAddTree08Surface:         q.Table15.FixedCropAddTree08.Count.AsFloat(),
			FixedCropAddTree08Name:            q.Table15.FixedCropAddTree08.AddCrop.CropName,
			FixedCropAddTree08Count:           int(q.Table15.FixedCropAddTree08.Count.TreeCount),
			FixedCropAddTree08Id:              int(q.Table15.FixedCropAddTree08.AddCrop.Id),
			CuttingMachineDiesel:              q.Table16.CuttingMachine.DieselInt(),
			CuttingMachineElectric:            q.Table16.CuttingMachine.ElectricInt(),
			CuttingMachineManual:              q.Table16.CuttingMachine.ManualInt(),
			FertilizerDiesel:                  q.Table16.Fertilizer.DieselInt(),
			FertilizerElectric:                q.Table16.Fertilizer.ElectricInt(),
			FertilizerManual:                  q.Table16.Fertilizer.ManualInt(),
			FishingNetDiesel:                  q.Table16.FishingNet.DieselInt(),
			FishingNetElectric:                q.Table16.FishingNet.ElectricInt(),
			FishingNetManual:                  q.Table16.FishingNet.ManualInt(),
			HarvestMachineDiesel:              q.Table16.HarvestMachine.DieselInt(),
			HarvestMachineElectric:            q.Table16.HarvestMachine.ElectricInt(),
			HarvestMachineManual:              q.Table16.HarvestMachine.ManualInt(),
			LinerDiesel:                       q.Table16.Liner.DieselInt(),
			LinerElectric:                     q.Table16.Liner.ElectricInt(),
			LinerManual:                       q.Table16.Liner.ManualInt(),
			PlowDiesel:                        q.Table16.Plow.DieselInt(),
			PlowElectric:                      q.Table16.Plow.ElectricInt(),
			PlowManual:                        q.Table16.Plow.ManualInt(),
			PowerPumpDiesel:                   q.Table16.PowerPump.DieselInt(),
			PowerPumpElectric:                 q.Table16.PowerPump.ElectricInt(),
			PowerPumpManual:                   q.Table16.PowerPump.ManualInt(),
			PowerTrillerDiesel:                q.Table16.PowerTriller.DieselInt(),
			PowerTrillerElectric:              q.Table16.PowerTriller.ElectricInt(),
			PowerTrillerManual:                q.Table16.PowerTriller.ManualInt(),
			SowingMachineDiesel:               q.Table16.SowingMachine.DieselInt(),
			SowingMachineElectric:             q.Table16.SowingMachine.ElectricInt(),
			SowingMachineManual:               q.Table16.SowingMachine.ManualInt(),
			TractorDiesel:                     q.Table16.Tractor.DieselInt(),
			TractorElectric:                   q.Table16.Tractor.ElectricInt(),
			TractorManual:                     q.Table16.Tractor.ManualInt(),
			TubeWellDiesel:                    q.Table16.TubeWell.DieselInt(),
			TubeWellElectric:                  q.Table16.TubeWell.ElectricInt(),
			TubeWellManual:                    q.Table16.TubeWell.ManualInt(),
			PeopleAbove15Female:               int(q.Table9.PeopleAbove15.Count.Female),
			PeopleAbove15Hijra:                int(q.Table9.PeopleAbove15.Count.Hijra),
			PeopleAbove15Male:                 int(q.Table9.PeopleAbove15.Count.Male),
			PeopleBetween10to14Female:         int(q.Table9.PeopleBetween10to14.Count.Female),
			PeopleBetween10to14Hijra:          int(q.Table9.PeopleBetween10to14.Count.Hijra),
			PeopleBetween10to14Male:           int(q.Table9.PeopleBetween10to14.Count.Male),
			TotalMemberFemale:                 int(q.Table9.TotalMember.Count.Female),
			TotalMemberHijra:                  int(q.Table9.TotalMember.Count.Hijra),
			TotalMemberMale:                   int(q.Table9.TotalMember.Count.Male),
		})

	}
	root = Root{
		CollectionName: json.CollectionName,
		Questionnaire:  questionnaires,
		TallySheet: &TallySheet{
			TallySheetBarcode:            json.TallySheet.TallySheetBarcode,
			TallySheetNo:                 json.TallySheet.TallySheetNo,
			FormName:                     json.TallySheet.FormName,
			OriginalAgriProfessionals:    json.TallySheet.IncludedInCountingArea.AgriProfessionals,
			OriginalFishingProfessionals: json.TallySheet.IncludedInCountingArea.FishingProfessionals,
			OriginalHouse5More:           json.TallySheet.IncludedInCountingArea.House5More,
			OriginalHouseFisheries:       json.TallySheet.IncludedInCountingArea.HouseFisheries,
			OriginalHouseNoLand:          json.TallySheet.IncludedInCountingArea.HouseNoLand,
			OriginalHouseReceivedLand:    json.TallySheet.IncludedInCountingArea.HouseReceivedLand,
			OriginalTotalHouse:           json.TallySheet.IncludedInCountingArea.TotalHouse,
			DataEditorName:               json.TallySheet.InformationDataCollector.DataEditorName,
			DataEditorSignature:          json.TallySheet.InformationDataCollector.DataEditorSignature,
			DataEntryName:                json.TallySheet.InformationDataCollector.DataEntryName,
			DataEntrySignature:           json.TallySheet.InformationDataCollector.DataEntrySignature,
			EnumDay:                      json.TallySheet.InformationDataCollector.EnumDay,
			EnumMobileNumber:             json.TallySheet.InformationDataCollector.EnumMobileNumber,
			EnumMonth:                    json.TallySheet.InformationDataCollector.EnumMonth,
			EnumName:                     json.TallySheet.InformationDataCollector.EnumName,
			EnumSignature:                json.TallySheet.InformationDataCollector.EnumSignature,
			EnumYear:                     json.TallySheet.InformationDataCollector.EnumYear,
			SupervisorDay:                json.TallySheet.InformationDataCollector.SupervisorDay,
			SupervisorMobileNumber:       json.TallySheet.InformationDataCollector.SupervisorMobileNumber,
			SupervisorMonth:              json.TallySheet.InformationDataCollector.SupervisorMonth,
			SupervisorName:               json.TallySheet.InformationDataCollector.SupervisorName,
			SupervisorSignature:          json.TallySheet.InformationDataCollector.SupervisorSignature,
			SupervisorYear:               json.TallySheet.InformationDataCollector.SupervisorYear,
			ZonalDay:                     json.TallySheet.InformationDataCollector.ZonalDay,
			ZonalMobileNumber:            json.TallySheet.InformationDataCollector.ZonalMobileNumber,
			ZonalMonth:                   json.TallySheet.InformationDataCollector.ZonalMonth,
			ZonalName:                    json.TallySheet.InformationDataCollector.ZonalName,
			ZonalSignature:               json.TallySheet.InformationDataCollector.ZonalSignature,
			ZonalYear:                    json.TallySheet.InformationDataCollector.ZonalYear,
			GeocodeID: fmt.Sprintf("%02d.%02d.%03d.%03d.%03d.%d",
				json.TallySheet.InformationOfTheCountingArea.District.Code,
				json.TallySheet.InformationOfTheCountingArea.Upazila.Code,
				json.TallySheet.InformationOfTheCountingArea.Union.Code,
				json.TallySheet.InformationOfTheCountingArea.Mouza.Code,
				json.TallySheet.InformationOfTheCountingArea.CountingArea.Code,
				json.TallySheet.InformationOfTheCountingArea.RMO.Code),
			OriginalBuffaloCount: json.TallySheet.SituatedAtTheCountingArea.BuffaloCount,
			OriginalCockCount:    json.TallySheet.SituatedAtTheCountingArea.CockCount,
			OriginalCowCount:     json.TallySheet.SituatedAtTheCountingArea.CowCount,
			OriginalDuckCount:    json.TallySheet.SituatedAtTheCountingArea.DuckCount,
			OriginalGoatCount:    json.TallySheet.SituatedAtTheCountingArea.GoatCount,
			OriginalSheepCount:   json.TallySheet.SituatedAtTheCountingArea.SheepCount,
			OriginalTurkeyCount:  json.TallySheet.SituatedAtTheCountingArea.TurkeyCount,
			Questionnaires:       questionnaires,
		},
	}

	return
}

type TallySheet struct {
	TallySheetBarcode            string
	TallySheetNo                 string `pg:",pk"`
	FormName                     string
	OriginalAgriProfessionals    int
	UpdatedAgriProfessionals     *int // OriginalAgriProfessionals if no update, count of Questionnaire.AgriLaborCode != 0
	OriginalFishingProfessionals int
	UpdatedFishingProfessionals  *int // OriginalFishingProfessionals if no update, count of Questionnaire.IsFishingRelated != 0
	OriginalHouse5More           int
	UpdatedHouse5More            *int // OriginalHouse5More if no update, sum of Questionnaire.CultivatedLand > 0.05
	OriginalHouseFisheries       int
	UpdatedHouseFisheries        *int // OriginalHouseFisheries if no update, sum of Questionnaire.PermanentUnusedLand
	OriginalHouseNoLand          int
	UpdatedHouseNoLand           *int // HouseNoLand if no update, sum of Questionnaire.TotalLand
	OriginalHouseReceivedLand    int
	UpdatedHouseReceivedLand     *int // OriginalHouseReceivedLand if no update, sum of Questionnaire.CultivatedLand
	OriginalTotalHouse           int
	UpdatedTotalHouse            *int // OriginalTotalHouse if no update, count of Questionnaire.HouseSerial is not null
	DataEditorName               string
	DataEditorSignature          int
	DataEntryName                string
	DataEntrySignature           int
	EnumDay                      int
	EnumMobileNumber             string
	EnumMonth                    int
	EnumName                     string
	EnumSignature                int
	EnumYear                     int
	SupervisorDay                int
	SupervisorMobileNumber       string
	SupervisorMonth              int
	SupervisorName               string
	SupervisorSignature          int
	SupervisorYear               string
	ZonalDay                     int
	ZonalMobileNumber            string
	ZonalMonth                   int
	ZonalName                    string
	ZonalSignature               int
	ZonalYear                    string
	GeocodeID                    string `pg:",notnull,type:ltree"` // InformationOfTheCountingArea in original JSON
	//GeoCode                      *ag.GeoCodes `pg:"rel:has-one,fk:geocode_id"`
	OriginalBuffaloCount int
	UpdatedBuffaloCount  *int // OriginalBuffaloCount if no update, sum of Questionnaire.BuffaloAtHome + Questionnaire.BuffaloAtFarm
	OriginalCockCount    int
	UpdatedCockCount     *int // OriginalCockCount if no update, sum of Questionnaire.CockAtHome + Questionnaire.CockAtFarm
	OriginalCowCount     int
	UpdatedCowCount      *int // OriginalCowCount if no update, sum of Questionnaire.CowAtHome + Questionnaire.CowAtFarm
	OriginalDuckCount    int
	UpdatedDuckCount     *int // OriginalDuckCount if no update, sum of Questionnaire.DuckAtHome Questionnaire.DuckAtFarm
	OriginalGoatCount    int
	UpdatedGoatCount     *int // OriginalGoatCount if no update, sum of Questionnaire.GoatAtHome + Questionnaire.GoatAtFarm
	OriginalSheepCount   int
	UpdatedSheepCount    *int // OriginalSheepCount if no update, sum of Questionnaire.SheepAtHome + Questionnaire.SheepAtFarm
	OriginalTurkeyCount  int
	UpdatedTurkeyCount   *int             // OriginalTurkeyCount if no update, sum of Questionnaire.TurkeyAtHome + Questionnaire.TurkeyAtFarm
	Questionnaires       []*Questionnaire `pg:"rel:has-many"`
}

type Questionnaire struct {
	QuestionnaireNum   string `pg:",pk"`
	QuestionnaireEmpty bool
	FormName           string
	BookletNumber      string
	//Booklet                           *ag.Booklet `pg:"rel:has-one, on_update: CASCADE"`
	TallySheetNo     string      // -> TallySheet.TallySheetNo
	TallySheet       *TallySheet `pg:"rel:has-one, on_update: CASCADE"`
	AgriLaborCode    int
	FatherName       string
	OwnerGender      int
	OwnerName        string
	HouseSerial      int
	IsFishingRelated int
	MobileNo         string
	OwnerEduCode     int `pg:",use_zero"`
	OwnerProfesCode  int
	IsUnion          int
	GeocodeID        string `pg:",notnull,type:ltree"` // InformationOfTheCountingArea in original JSON
	//GeoCode                           *ag.GeoCodes `pg:"rel:has-one,fk:geocode_id"`
	Village                           int // Location.village in questionnaire
	BusinessEntityLand                float64
	CultivatedLand                    float64
	IrrigationLand                    float64
	LandGiven                         float64
	LandTaken                         float64
	NurseryLand                       float64
	OperatingLand                     float64
	PermanentCultivatedLand           float64
	PermanentUnusedLand               float64
	PondLandNotBlank                  float64
	RecentUnusedLand                  float64
	ResidenceLand                     float64
	SaltCultivationLand               float64
	SinkChannelLand                   float64
	TempCultivatedLand                float64
	TotalLand                         float64
	UncultivatedLand                  float64
	CreekLand                         float64
	FishCageCultiveLand               float64
	FishCultivationLand               float64
	FishSaltCultiveLand               float64
	MixedCultivationLand              float64
	PaddyCultivationLand              float64
	PondLand                          float64
	CockAtFarm                        int
	CockAtHome                        int
	DuckAtFarm                        int
	DuckAtHome                        int
	PigeonAtFarm                      int
	PigeonAtHome                      int
	QuailAtFarm                       int
	QuailAtHome                       int
	TurkeyAtFarm                      int
	TurkeyAtHome                      int
	BuffaloAtFarm                     int
	BuffaloAtHome                     int
	CowAtFarm                         int
	CowAtHome                         int
	GoatAtFarm                        int
	GoatAtHome                        int
	HorseAtFarm                       int
	HorseAtHome                       int
	PigAtFarm                         int
	PigAtHome                         int
	SheepAtFarm                       int
	SheepAtHome                       int
	TempAddCrop01Name                 string
	TempAddCrop01Id                   int
	TempAddCrop01Surface              float64
	TempAddCrop02Name                 string
	TempAddCrop02Id                   int
	TempAddCrop02Surface              float64
	TempAddCrop03Name                 string
	TempAddCrop03Id                   int
	TempAddCrop03Surface              float64
	TempAddCrop04Name                 string
	TempAddCrop04Id                   int
	TempAddCrop04Surface              float64
	TempAddCrop05Name                 string
	TempAddCrop05Id                   int
	TempAddCrop05Surface              float64
	TempAddCrop06Name                 string
	TempAddCrop06Id                   int
	TempAddCrop06Surface              float64
	TempAddCrop07Name                 string
	TempAddCrop07Id                   int
	TempAddCrop07Surface              float64
	TempAddCrop08Name                 string
	TempAddCrop08Id                   int
	TempAddCrop08Surface              float64
	TempAddCrop09Name                 string
	TempAddCrop09Id                   int
	TempAddCrop09Surface              float64
	TempCrop101Aush                   float64
	TempCrop102Almond                 float64
	TempCrop103Boro                   float64
	TempCrop104Wheat                  float64
	TempCrop105Maize                  float64
	TempCrop112Lentil                 float64
	TempCrop113Safran                 float64
	TempCrop114Moog                   float64
	TempCrop121Potato                 float64
	TempCrop122Brinjal                float64
	TempCrop123Radish                 float64
	TempCrop124Bean                   float64
	TempCrop125Tomato                 float64
	TempCrop127Multitude              float64
	TempCrop128ladyfinger             float64
	TempCrop129Cucumber               float64
	TempCrop130BitterGourd            float64
	TempCrop131Ground                 float64
	TempCrop132Pumplin                float64
	TempCrop134Cauliflower            float64
	TempCrop135Cabbage                float64
	TempCrop157Onion                  float64
	TempCrop158Garlic                 float64
	TempCrop159Ginger                 float64
	TempCrop160Turmeric               float64
	TempCrop161Chili                  float64
	TempCrop167Mustard                float64
	TempCrop169Nuts                   float64
	TempCrop175Banana                 float64
	TempCrop176Papaya                 float64
	TempCrop177WMelon                 float64
	TempCrop179PineApple              float64
	TempCrop182Jute                   float64
	TempCrop185SugarCane              float64
	TempCrop203Seeded                 float64
	FixedCrop501MongoSurface          float64
	FixedCrop501MongoCount            int
	FixedCrop502BerrySurface          float64
	FixedCrop502BerryCount            int
	FixedCrop503JackFruitSurface      float64
	FixedCrop503JackFruitCount        int
	FixedCrop504LitchiSurface         float64
	FixedCrop504LitchiCount           int
	FixedCrop505GuavaSurface          float64
	FixedCrop505GuavaCount            int
	FixedCrop506CoconutSurface        float64
	FixedCrop506CoconutCount          int
	FixedCrop507PlumSurface           float64
	FixedCrop507PlumCount             int
	FixedCrop508HogPlumSurface        float64
	FixedCrop508HogPlumCount          int
	FixedCrop510DateSurface           float64
	FixedCrop510DateCount             int
	FixedCrop511PalmyraSurface        float64
	FixedCrop511PalmyraCount          int
	FixedCrop512BellSurface           float64
	FixedCrop512BellCount             int
	FixedCrop521AverrhoaSurface       float64
	FixedCrop521AverrhoaCount         int
	FixedCrop522TamarindSurface       float64
	FixedCrop522TamarindCount         int
	FixedCrop523LemonSurface          float64
	FixedCrop523LemonCount            int
	FixedCrop524GrapeSurface          float64
	FixedCrop524GrapeCount            int
	FixedCrop538BattleLeafSurface     float64
	FixedCrop538BattleLeafCount       int
	FixedCrop539ArecaCatechuSurface   float64
	FixedCrop539ArecaCatechuCount     int
	FixedCrop546BambooSurface         float64
	FixedCrop546BambooCount           int
	FixedCrop548HardwoodSurface       float64
	FixedCrop548HardwoodCount         int
	FixedCrop549RainTreeSurface       float64
	FixedCrop549RainTreeCount         int
	FixedCrop550MahoganySurface       float64
	FixedCrop550MahoganyCount         int
	FixedCrop551TectonaGrandisSurface float64
	FixedCrop551TectonaGrandisCount   int
	FixedCrop572NeemSurface           float64
	FixedCrop572NeemCount             int
	FixedCropAddTree01Surface         float64
	FixedCropAddTree01Name            string
	FixedCropAddTree01Count           int
	FixedCropAddTree01Id              int
	FixedCropAddTree02Surface         float64
	FixedCropAddTree02Name            string
	FixedCropAddTree02Count           int
	FixedCropAddTree02Id              int
	FixedCropAddTree03Surface         float64
	FixedCropAddTree03Name            string
	FixedCropAddTree03Count           int
	FixedCropAddTree03Id              int
	FixedCropAddTree04Surface         float64
	FixedCropAddTree04Name            string
	FixedCropAddTree04Count           int
	FixedCropAddTree04Id              int
	FixedCropAddTree05Surface         float64
	FixedCropAddTree05Name            string
	FixedCropAddTree05Count           int
	FixedCropAddTree05Id              int
	FixedCropAddTree06Surface         float64
	FixedCropAddTree06Name            string
	FixedCropAddTree06Count           int
	FixedCropAddTree06Id              int
	FixedCropAddTree07Surface         float64
	FixedCropAddTree07Name            string
	FixedCropAddTree07Count           int
	FixedCropAddTree07Id              int
	FixedCropAddTree08Surface         float64
	FixedCropAddTree08Name            string
	FixedCropAddTree08Count           int
	FixedCropAddTree08Id              int
	CuttingMachineDiesel              int
	CuttingMachineElectric            int
	CuttingMachineManual              int
	FertilizerDiesel                  int
	FertilizerElectric                int
	FertilizerManual                  int
	FishingNetDiesel                  int
	FishingNetElectric                int
	FishingNetManual                  int
	HarvestMachineDiesel              int
	HarvestMachineElectric            int
	HarvestMachineManual              int
	LinerDiesel                       int
	LinerElectric                     int
	LinerManual                       int
	PlowDiesel                        int
	PlowElectric                      int
	PlowManual                        int
	PowerPumpDiesel                   int
	PowerPumpElectric                 int
	PowerPumpManual                   int
	PowerTrillerDiesel                int
	PowerTrillerElectric              int
	PowerTrillerManual                int
	SowingMachineDiesel               int
	SowingMachineElectric             int
	SowingMachineManual               int
	TractorDiesel                     int
	TractorElectric                   int
	TractorManual                     int
	TubeWellDiesel                    int
	TubeWellElectric                  int
	TubeWellManual                    int
	PeopleAbove15Female               int
	PeopleAbove15Hijra                int
	PeopleAbove15Male                 int
	PeopleBetween10to14Female         int
	PeopleBetween10to14Hijra          int
	PeopleBetween10to14Male           int
	TotalMemberFemale                 int
	TotalMemberHijra                  int
	TotalMemberMale                   int
}

type OverwrittenQuestionnaire struct {
	UpdateDate time.Time `pg:",pk"`
	UpdatedBy  string
	Questionnaire
}
