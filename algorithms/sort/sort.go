package sort

import "algorithms/algorithms/myutil"

func Sort() {
	myutil.Rival(10000, SeSort)
	myutil.Rival(10000, BSort)
	myutil.Rival(10000, CombSort)
	myutil.Rival(10000, CocktailSort)
	myutil.Rival(10000, CeSort)
	myutil.Rival(10000, QSort)
	myutil.Rival(10000, MSort)
	myutil.Rival(10000, SeSort)

}
