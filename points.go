package main

import (
    "math"
    "strconv"
    "strings"
    "time"
)

func calculatePoints(receipt Receipt) int {
    points := 0

    points += len(alphanumericChars(receipt.Retailer))

    if isWholeDollar(receipt.Total) {
        points += 50
    }

    if isMultipleOfQuarter(receipt.Total) {
        points += 25
    }

    points += (len(receipt.Items) / 2) * 5

    for _, item := range receipt.Items {
        if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
            itemPrice, _ := strconv.ParseFloat(item.Price, 64)
            points += int(math.Ceil(itemPrice * 0.2))
        }
    }

    purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
    if purchaseDate.Day()%2 != 0 {
        points += 6
    }

    purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
    if purchaseTime.Hour() == 14 {
        points += 10
    }

    return points
}

func alphanumericChars(s string) string {
    result := ""
    for _, ch := range s {
        if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
            result += string(ch)
        }
    }
    return result
}

func isWholeDollar(total string) bool {
    totalAmount, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return false
    }
    return totalAmount == math.Floor(totalAmount)
}

func isMultipleOfQuarter(total string) bool {
    totalAmount, err := strconv.ParseFloat(total, 64)
    if err != nil {
        return false
    }
    return math.Mod(totalAmount, 0.25) == 0
}
