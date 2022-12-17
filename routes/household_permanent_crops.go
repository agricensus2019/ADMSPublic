package routes

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func (srv *Server) FormatHouseholdPermanentCrops(division, district, upazilla, union, mouza string, q *searchQuery, geoLocation string) (tableAndDonut string, err error) {
	c, err := srv.Db.GetPermanantCrops(division, district, upazilla, union, mouza)
	if err != nil {
		return "", err
	}

	p := message.NewPrinter(language.English)

	tableAndDonut = fmt.Sprintf(`
	<div class="x_content">
	<h4>Result<small> </small></h4>
	<h5>Table name : %s</h5>
	<h5>Number of farm holdings : %s</h5>
	<h5>Crop area : %s</h5>
	<table id="datatable-buttons" class="table table-striped">
		<thead>
			<tr>
				<th>Crop code and name</th>
				<th>Total Permanent crop area (acres)</th>
				<th>Percentage</th>
			</tr>
		</thead>
		<tbody>
		<tr> <td><b>501 Mango</b> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>502 Berry</b></td> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>503 Jackfruit</b></td> <td>%s</td> <td>%s</td> </tr>          
		<tr> <td><b>504 Litchi</b></td> <td>%s</td> <td>%s</td> </tr>         
		<tr> <td><b>505 Guava</b></td> <td>%s</td> <td>%s</td> </tr>         
		<tr> <td><b>506 Coconut</b></td> <td>%s</td> <td>%s</td> </tr>
		<tr> <td><b>507 Jujube</b></td> <td>%s</td> <td>%s</td> </tr> 
		<tr> <td><b>508 Hog Plum</b></td> <td>%s</td> <td>%s</td> </tr>  
		<tr> <td><b>509 Olive</b></td> <td>%s</td> <td>%s</td> </tr>  
		<tr> <td><b>510 Date Palm</b></td> <td>%s</td> <td>%s</td> </tr>    
		<tr> <td><b>511 Palm</b></td> <td>%s</td> <td>%s</td> </tr>   
		<tr> <td><b>512 Wood Apple</b></td> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>513 Limonia Acidissima</b></td> <td>%s</td> <td>%s</td> </tr> 
		<tr> <td><b>514 Star Apple</b></td> <td>%s</td> <td>%s</td> </tr>    
		<tr> <td><b>515 Carissa Carandas</b></td> <td>%s</td> <td>%s</td> </tr>    
		<tr> <td><b>516 Sugar Apple</b></td> <td>%s</td> <td>%s</td> </tr>           
		<tr> <td><b>517 Custard Apple</b></td> <td>%s</td> <td>%s</td> </tr>      
		<tr> <td><b>518 Pomegranate</b></td> <td>%s</td> <td>%s</td> </tr>       
		<tr> <td><b>519 Sapodilla</b></td> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>520 Monkey Jack</b></td> <td>%s</td> <td>%s</td> </tr>   
		<tr> <td><b>521 Carambola</b></td> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>522 Tamarind</b></td> <td>%s</td> <td>%s</td> </tr>       
		<tr> <td><b>523 Lemon</b></td> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>524 Pomelo</b></td> <td>%s</td> <td>%s</td> </tr>          
		<tr> <td><b>525 Gosseberry</b></td> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>526 Baccaurea Motleyana</b></td> <td>%s</td> <td>%s</td> </tr>     
		<tr> <td><b>527 Otaheite Gooseberry</b></td> <td>%s</td> <td>%s</td> </tr>     
		<tr> <td><b>528 Elephant Apple</b></td> <td>%s</td> <td>%s</td> </tr> 
		<tr> <td><b>529 Orange</b></td> <td>%s</td> <td>%s</td> </tr>      
		<tr> <td><b>530 Citrus Macroptera</b></td> <td>%s</td> <td>%s</td> </tr>
		<tr> <td><b>531 Citrus Sinensis</b></td> <td>%s</td> <td>%s</td> </tr>                 
		<tr> <td><b>532 Bilimb</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>533 Persimmon</b></td> <td>%s</td> <td>%s</td> </tr>          
		<tr> <td><b>534 Figs</b></td> <td>%s</td> <td>%s</td> </tr>           
		<tr> <td><b>535 Dragon</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>536 Rambutan</b></td> <td>%s</td> <td>%s</td> </tr>              
		<!-- <tr> <td><b>537 Other Fruits</b></td> <td>%s</td> <td>%s</td> </tr>  -->             
		<tr> <td><b>538 Battle Leaf</b></td> <td>%s</td> <td>%s</td> </tr>          
		<tr> <td><b>539 Betel-Nut</b></td> <td>%s</td> <td>%s</td> </tr>                 
		<tr> <td><b>540 Tea</b></td> <td>%s</td> <td>%s</td> </tr>                  
		<!-- <tr> <td><b>541 Other Addictives/Other Drugs and Narcotics</b></td> <td>%s</td> <td>%s</td> </tr> -->      
		<tr> <td><b>542 Bay Leaf</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>543 Cinnamon</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>544 Cardamom</b></td> <td>%s</td> <td>%s</td> </tr>              
		<!-- <tr> <td><b>545 Other Spices</b></td> <td>%s</td> <td>%s</td> </tr> -->                
		<tr> <td><b>546 Bamboo</b></td> <td>%s</td> <td>%s</td> </tr>                 
		<tr> <td><b>547 Cane</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>548 Lebbek</b></td> <td>%s</td> <td>%s</td> </tr>          
		<tr> <td><b>549 Rain Tree</b></td> <td>%s</td> <td>%s</td> </tr>              
		<tr> <td><b>550 Mahogany</b></td> <td>%s</td> <td>%s</td> </tr>      
		<tr> <td><b>551 Teak</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>552 Auri</b></td> <td>%s</td> <td>%s</td> </tr>        
		<tr> <td><b>553 Eucalyptus</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>554 Dalbergia Sissoo</b></td> <td>%s</td> <td>%s</td> </tr>           
		<tr> <td><b>555 Gmelina Arborea</b></td> <td>%s</td> <td>%s</td> </tr>          
		<tr> <td><b>556 Rhizophora Apiculata</b></td> <td>%s</td> <td>%s</td> </tr>           
		<tr> <td><b>557 Banyan</b></td> <td>%s</td> <td>%s</td> </tr>                 
		<tr> <td><b>559 Typha</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>560 Polyalthia Longifolia</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>561 AIstonia Scholaris</b></td> <td>%s</td> <td>%s</td> </tr>              
		<tr> <td><b>562 Jarul</b></td> <td>%s</td> <td>%s</td> </tr>                 
		<tr> <td><b>563 Gum Arabic</b></td> <td>%s</td> <td>%s</td> </tr>             
		<tr> <td><b>564 Monkey Jack</b></td> <td>%s</td> <td>%s</td> </tr>           
		<tr> <td><b>565 Golden Shower</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>566 Rubber</b></td> <td>%s</td> <td>%s</td> </tr>                 
		<tr> <td><b>567 Aquilaria Malaccensis</b></td> <td>%s</td> <td>%s</td> </tr>
		<tr> <td><b>568 Sakhua or Shala</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>569 Artocarpus Chama</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>570 Albizia Lebbeck</b></td> <td>%s</td> <td>%s</td> </tr>                  
		<!-- <tr> <td><b>571 Other Wood and Forestry</b></td> <td>%s</td> <td>%s</td> </tr> -->                
		<tr> <td><b>572 Neem</b></td> <td>%s</td> <td>%s</td> </tr>               
		<tr> <td><b>573 Arjun Tree</b></td> <td>%s</td> <td>%s</td> </tr>             
		<tr> <td><b>574 Bahera or Beleric</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>575 Myrobalan</b></td> <td>%s</td> <td>%s</td> </tr>       
		<tr> <td><b>577 Drumstick tree</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>579 Rose</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>580 Bakul</b></td> <td>%s</td> <td>%s</td> </tr>           
		<tr> <td><b>581 Neolamarckia Cadamba</b></td> <td>%s</td> <td>%s</td> </tr>                
		<tr> <td><b>582 Royal Poinciana</b></td> <td>%s</td> <td>%s</td> </tr>            
		<tr> <td><b>584 Silk-Cotton</b></td> <td>%s</td> <td>%s</td> </tr>          
		<tr> <td><b>585 Mulberry</b></td> <td>%s</td> <td>%s</td> </tr> 
						
		</tbody>
	</table>
	</div>
	<h7>Source: Bangladesh Bureau of Statistics. Report produced by Agriculture (Crops, Fisheries and Livestock) Census 2018 Project.</h7>
	`,
		fmt.Sprintf("%s <br> %s", getTableGenerationName(q.TableNumber), geoLocation),
		p.Sprintf("%d", c.NumberOfFarmHoldings),
		p.Sprintf("%.2f", c.CropArea),
		FormatFloat(c.P501a, 2), c.PercentageOfPermantCropArea("P501a"),
		FormatFloat(c.P502a, 2), c.PercentageOfPermantCropArea("P502a"),
		FormatFloat(c.P503a, 2), c.PercentageOfPermantCropArea("P503a"),
		FormatFloat(c.P504a, 2), c.PercentageOfPermantCropArea("P504a"),
		FormatFloat(c.P505a, 2), c.PercentageOfPermantCropArea("P505a"),
		FormatFloat(c.P506a, 2), c.PercentageOfPermantCropArea("P506a"),
		FormatFloat(c.P507a, 2), c.PercentageOfPermantCropArea("P507a"),
		FormatFloat(c.P508a, 2), c.PercentageOfPermantCropArea("P508a"),
		FormatFloat(c.P509a, 2), c.PercentageOfPermantCropArea("P509a"),
		FormatFloat(c.P510a, 2), c.PercentageOfPermantCropArea("P510a"),
		FormatFloat(c.P511a, 2), c.PercentageOfPermantCropArea("P511a"),
		FormatFloat(c.P512a, 2), c.PercentageOfPermantCropArea("P512a"),
		FormatFloat(c.P513a, 2), c.PercentageOfPermantCropArea("P513a"),
		FormatFloat(c.P514a, 2), c.PercentageOfPermantCropArea("P514a"),
		FormatFloat(c.P515a, 2), c.PercentageOfPermantCropArea("P515a"),
		FormatFloat(c.P516a, 2), c.PercentageOfPermantCropArea("P516a"),
		FormatFloat(c.P517a, 2), c.PercentageOfPermantCropArea("P517a"),
		FormatFloat(c.P518a, 2), c.PercentageOfPermantCropArea("P518a"),
		FormatFloat(c.P519a, 2), c.PercentageOfPermantCropArea("P519a"),
		FormatFloat(c.P520a, 2), c.PercentageOfPermantCropArea("P520a"),
		FormatFloat(c.P521a, 2), c.PercentageOfPermantCropArea("P521a"),
		FormatFloat(c.P522a, 2), c.PercentageOfPermantCropArea("P522a"),
		FormatFloat(c.P523a, 2), c.PercentageOfPermantCropArea("P523a"),
		FormatFloat(c.P524a, 2), c.PercentageOfPermantCropArea("P524a"),
		FormatFloat(c.P525a, 2), c.PercentageOfPermantCropArea("P525a"),
		FormatFloat(c.P526a, 2), c.PercentageOfPermantCropArea("P526a"),
		FormatFloat(c.P527a, 2), c.PercentageOfPermantCropArea("P527a"),
		FormatFloat(c.P528a, 2), c.PercentageOfPermantCropArea("P528a"),
		FormatFloat(c.P529a, 2), c.PercentageOfPermantCropArea("P529a"),
		FormatFloat(c.P530a, 2), c.PercentageOfPermantCropArea("P530a"),
		FormatFloat(c.P531a, 2), c.PercentageOfPermantCropArea("P531a"),
		FormatFloat(c.P532a, 2), c.PercentageOfPermantCropArea("P532a"),
		FormatFloat(c.P533a, 2), c.PercentageOfPermantCropArea("P533a"),
		FormatFloat(c.P534a, 2), c.PercentageOfPermantCropArea("P534a"),
		FormatFloat(c.P535a, 2), c.PercentageOfPermantCropArea("P535a"),
		FormatFloat(c.P536a, 2), c.PercentageOfPermantCropArea("P536a"),
		// FormatFloat(c.P537a, 2), c.PercentageOfPermantCropArea("P537a"),
		FormatFloat(c.P538a, 2), c.PercentageOfPermantCropArea("P538a"),
		FormatFloat(c.P539a, 2), c.PercentageOfPermantCropArea("P539a"),
		FormatFloat(c.P540a, 2), c.PercentageOfPermantCropArea("P540a"),
		// FormatFloat(c.P541a, 2), c.PercentageOfPermantCropArea("P541a"),
		FormatFloat(c.P542a, 2), c.PercentageOfPermantCropArea("P542a"),
		FormatFloat(c.P543a, 2), c.PercentageOfPermantCropArea("P543a"),
		FormatFloat(c.P544a, 2), c.PercentageOfPermantCropArea("P544a"),
		// FormatFloat(c.P545a, 2), c.PercentageOfPermantCropArea("P545a"),
		FormatFloat(c.P546a, 2), c.PercentageOfPermantCropArea("P546a"),
		FormatFloat(c.P547a, 2), c.PercentageOfPermantCropArea("P547a"),
		FormatFloat(c.P548a, 2), c.PercentageOfPermantCropArea("P548a"),
		FormatFloat(c.P549a, 2), c.PercentageOfPermantCropArea("P549a"),
		FormatFloat(c.P550a, 2), c.PercentageOfPermantCropArea("P550a"),
		FormatFloat(c.P551a, 2), c.PercentageOfPermantCropArea("P551a"),
		FormatFloat(c.P552a, 2), c.PercentageOfPermantCropArea("P552a"),
		FormatFloat(c.P553a, 2), c.PercentageOfPermantCropArea("P553a"),
		FormatFloat(c.P554a, 2), c.PercentageOfPermantCropArea("P554a"),
		FormatFloat(c.P555a, 2), c.PercentageOfPermantCropArea("P555a"),
		FormatFloat(c.P556a, 2), c.PercentageOfPermantCropArea("P556a"),
		FormatFloat(c.P557a, 2), c.PercentageOfPermantCropArea("P557a"),
		FormatFloat(c.P559a, 2), c.PercentageOfPermantCropArea("P559a"),
		FormatFloat(c.P560a, 2), c.PercentageOfPermantCropArea("P560a"),
		FormatFloat(c.P561a, 2), c.PercentageOfPermantCropArea("P561a"),
		FormatFloat(c.P562a, 2), c.PercentageOfPermantCropArea("P562a"),
		FormatFloat(c.P563a, 2), c.PercentageOfPermantCropArea("P563a"),
		FormatFloat(c.P564a, 2), c.PercentageOfPermantCropArea("P564a"),
		FormatFloat(c.P565a, 2), c.PercentageOfPermantCropArea("P565a"),
		FormatFloat(c.P566a, 2), c.PercentageOfPermantCropArea("P566a"),
		FormatFloat(c.P567a, 2), c.PercentageOfPermantCropArea("P567a"),
		FormatFloat(c.P568a, 2), c.PercentageOfPermantCropArea("P568a"),
		FormatFloat(c.P569a, 2), c.PercentageOfPermantCropArea("P569a"),
		FormatFloat(c.P570a, 2), c.PercentageOfPermantCropArea("P570a"),
		// FormatFloat(c.P571a, 2), c.PercentageOfPermantCropArea("P571a"),
		FormatFloat(c.P572a, 2), c.PercentageOfPermantCropArea("P572a"),
		FormatFloat(c.P573a, 2), c.PercentageOfPermantCropArea("P573a"),
		FormatFloat(c.P574a, 2), c.PercentageOfPermantCropArea("P574a"),
		FormatFloat(c.P575a, 2), c.PercentageOfPermantCropArea("P575a"),
		FormatFloat(c.P577a, 2), c.PercentageOfPermantCropArea("P577a"),
		FormatFloat(c.P579a, 2), c.PercentageOfPermantCropArea("P579a"),
		FormatFloat(c.P580a, 2), c.PercentageOfPermantCropArea("P580a"),
		FormatFloat(c.P581a, 2), c.PercentageOfPermantCropArea("P581a"),
		FormatFloat(c.P582a, 2), c.PercentageOfPermantCropArea("P582a"),
		FormatFloat(c.P584a, 2), c.PercentageOfPermantCropArea("P584a"),
		FormatFloat(c.P585a, 2), c.PercentageOfPermantCropArea("P585a"),
	)

	return
}
