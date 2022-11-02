package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind=="car" || kind=="truck"{
        return true
    }else{
        return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    if len(option1)>=len(option2){
        for i:=0;i<len(option2);i++{
            if option1[i]>option2[i]{
                return option2 + " is clearly the better choice."
            }
        	if option1[i]<option2[i]{
                return option1 + " is clearly the better choice."
            }
        }
        return option2 + " is clearly the better choice."
    }else{
        for i:=0;i<len(option1);i++{
            if option1[i]>option2[i]{
                return option2 + " is clearly the better choice."
            }
        	if option1[i]<option2[i]{
                return option1 + " is clearly the better choice."
            }
        }
    	return option1 + " is clearly the better choice."
    }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <3{
        return originalPrice*0.8
    }else if age>=10{
    	return originalPrice*0.5
    }else{
    	return originalPrice*0.7
    }
}
