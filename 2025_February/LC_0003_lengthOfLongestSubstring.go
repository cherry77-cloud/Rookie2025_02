func lengthOfLongestSubstring(s string) int {
    // 用来记录每个字符出现的位置
    charIndexMap := make(map[byte]int)
    maxLen := 0
    left := 0  // 滑动窗口左边界

    for right := 0; right < len(s); right++ {
        // 当前字符
        ch := s[right]
        // 如果当前字符在窗口内出现过，则更新左边界
        if index, found := charIndexMap[ch]; found && index >= left {
            left = index + 1
        }
        // 更新当前字符的位置
        charIndexMap[ch] = right
        // 计算窗口大小，更新最大长度
        if currentLen := right - left + 1; currentLen > maxLen {
            maxLen = currentLen
        }
    }
    
    return maxLen
}
