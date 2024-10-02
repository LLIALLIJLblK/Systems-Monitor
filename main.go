package main

import(
	"log"
	"GO_LANG/handlers/monitoring"
)


func main(){

	metrics, err := monitoring.GetSysMetrics()
	if err != nil {
		log.Fatal(err)
	}

	monitoring.PrintSystemMetrics(metrics)
}


//---------------------------------------------
// написать такую хуйнию чтоб она мониторла все системы ppa zina 
//---------------------------------------------

// Total: Общий объем физической памяти.
// Available: Объем доступной памяти.
// Used: Объем используемой памяти.
// UsedPercent: Процент используемой памяти.
// Free: Объем свободной памяти.
// Active: Объем активной памяти.
// Inactive: Объем неактивной памяти.
// Wired: Объем зарезервированной памяти.


// {"path":"/"
// ,"fstype":"apfs"
// ,"total":494384795648
// ,"free":342344876032
// ,"used":152039919616
// ,"usedPercent":30.75335668782416
// ,"inodesTotal":3343617374
// ,"inodesUsed":405694
// ,"inodesFree":3343211680
// ,"inodesUsedPercent":0.012133385929702374}