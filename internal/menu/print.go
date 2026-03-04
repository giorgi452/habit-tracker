/* Package menu */
package menu

import "fmt"

func Print() (int8, error) {
	var operation int8
	fmt.Println("--- Welcome to Habit Tracker ---")
	fmt.Println("[1] Add Habit ")
	fmt.Println("[2] Start Habit")
	fmt.Println("[3] Edit Habit")
	fmt.Println("[4] Delete Habit")
	fmt.Println("[5] Exit")
	fmt.Print("Choose operation: ")

	_, err := fmt.Scanln(&operation)
	if err != nil {
		var dump string
		fmt.Scan(&dump)

		return 0, fmt.Errorf("please enter a number, not text")
	}

	if operation < 1 || operation > 5 {
		fmt.Println("⚠️ Error: Choice must be between 1 and 5.")
		return Print()
	}
	return operation, nil
}
