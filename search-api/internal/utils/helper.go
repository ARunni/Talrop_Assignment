package utils

import (
	"regexp"
	"search-api/internal/domain"
	"search-api/internal/utils/interfaces"
	"sort"
	"strings"
)



type helper struct{}

func NewHelper() interfaces.Helper {
	return &helper{}
}

// Rank users based on phonetic similarity
func (r *helper) RankUsers(inputName string, users []domain.User) []domain.User {
	for i := range users {
		users[i].Score = computePhoneticScore(inputName, users[i].Name)
	}

	// Sort users by score (descending)
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})

	return users
}
func (r *helper) IsAlphabetic(input string) bool {
	// Regular expression that allows only alphabetic characters (A-Z, a-z)
	re := regexp.MustCompile("^[a-zA-Z]+$")
	return re.MatchString(input)
}

func soundex(input string) string {
	if input == "" {
		return ""
	}

	// Step 1: Keep the first letter
	firstLetter := strings.ToUpper(string(input[0]))
	input = strings.ToLower(input)

	// Step 2: Map letters to Soundex digits
	replacements := map[rune]string{
		'b': "1", 'f': "1", 'p': "1", 'v': "1",
		'c': "2", 'g': "2", 'j': "2", 'k': "2", 'q': "2", 's': "2", 'x': "2", 'z': "2",
		'd': "3", 't': "3",
		'l': "4",
		'm': "5", 'n': "5",
		'r': "6",
	}

	// Step 3: Replace letters with digits
	var encoded strings.Builder
	for _, char := range input[1:] {
		if digit, exists := replacements[char]; exists {
			encoded.WriteString(digit)
		} else {
			encoded.WriteString("0") // Non-mapped letters are replaced with "0"
		}
	}

	// Step 4: Remove duplicate digits
	result := make([]rune, 0, len(encoded.String()))
	prev := rune(0)
	for _, char := range encoded.String() {
		if char != prev {
			result = append(result, char)
		}
		prev = char
	}

	// Step 5: Limit to 4 characters, padding with zeros if necessary
	soundexCode := firstLetter + string(result)
	soundexCode = strings.ReplaceAll(soundexCode, "0", "")
	if len(soundexCode) < 4 {
		soundexCode += strings.Repeat("0", 4-len(soundexCode))
	} else if len(soundexCode) > 4 {
		soundexCode = soundexCode[:4]
	}

	return soundexCode
}

// Compute score
func computePhoneticScore(input string, name string) float64 {
	inputKey := soundex(input)
	nameKey := soundex(name)

	// Normalize the edit distance to a score between 0.0 and 1.0
	distance := levenshtein(inputKey, nameKey)
	maxLen := max(len(inputKey), len(nameKey))
	score := 1.0 - float64(distance)/float64(maxLen)

	return score
}

// Finding the edit distance
func levenshtein(a, b string) int {
	// Create a 2D slice for dynamic programming
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the dp array
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Compute edit distance
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

// min
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

// max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
