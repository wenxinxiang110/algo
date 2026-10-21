# algo
Algorithm learning 



Directory：

```shell
.
├── README.md
├── dp
│ ├── package.go
│ └── package_test.go
├── go.mod
├── go.sum
├── kmp
│ ├── kmp.go
│ └── kmp_test.go
├── leetcode
│ ├── array
│ │ ├── duplicate.go
│ │ ├── easy.go
│ │ ├── intersection.go
│ │ ├── max_profit.go
│ │ ├── move_in_situ.go
│ │ ├── summaryRanges.go
│ │ ├── summaryRanges_test.go
│ │ ├── validMountainArray.go
│ │ └── validMountainArray_test.go
│ ├── dp
│ │ ├── climbStairs.go
│ │ └── maxSubArray.go
│ ├── editor
│ │ └── cn
│ │     ├── 1002_find-common-characters.go
│ │     ├── 1002_find-common-characters_test.go
│ │     ├── 101_symmetric-tree.go
│ │     ├── 102_binary-tree-level-order-traversal.go
│ │     ├── 1047_remove-all-adjacent-duplicates-in-string.go
│ │     ├── 1047_remove-all-adjacent-duplicates-in-string_test.go
│ │     ├── 1049_last-stone-weight-ii.go
│ │     ├── 1049_last-stone-weight-ii_test.go
│ │     ├── 104_maximum-depth-of-binary-tree.go
│ │     ├── 105_construct-binary-tree-from-preorder-and-inorder-traversal.go
│ │     ├── 105_construct-binary-tree-from-preorder-and-inorder-traversal_test.go
│ │     ├── 108_convert-sorted-array-to-binary-search-tree.go
│ │     ├── 108_convert-sorted-array-to-binary-search-tree_test.go
│ │     ├── 114_flatten-binary-tree-to-linked-list.go
│ │     ├── 114_flatten-binary-tree-to-linked-list_test.go
│ │     ├── 11_container-with-most-water.go
│ │     ├── 11_container-with-most-water_test.go
│ │     ├── 1209_remove-all-adjacent-duplicates-in-string-ii.go
│ │     ├── 124_binary-tree-maximum-path-sum.go
│ │     ├── 124_binary-tree-maximum-path-sum_test.go
│ │     ├── 128_longest-consecutive-sequence.go
│ │     ├── 128_longest-consecutive-sequence_test.go
│ │     ├── 131_palindrome-partitioning.go
│ │     ├── 131_palindrome-partitioning_test.go
│ │     ├── 136_single-number.py
│ │     ├── 138_copy-list-with-random-pointer.go
│ │     ├── 141_linked-list-cycle.go
│ │     ├── 142_linked-list-cycle-ii.go
│ │     ├── 144_binary-tree-preorder-traversal.go
│ │     ├── 146_lru-cache.go
│ │     ├── 146_lru-cache_test.go
│ │     ├── 148_sort-list.go
│ │     ├── 148_sort-list_test.go
│ │     ├── 151_reverse-words-in-a-string.go
│ │     ├── 151_reverse-words-in-a-string_test.go
│ │     ├── 155_min-stack.go
│ │     ├── 15_3sum.go
│ │     ├── 15_3sum_test.go
│ │     ├── 160_intersection-of-two-linked-lists.go
│ │     ├── 17_letter-combinations-of-a-phone-number.go
│ │     ├── 17_letter-combinations-of-a-phone-number_test.go
│ │     ├── 189_rotate-array.go
│ │     ├── 18_4sum.go
│ │     ├── 18_4sum_test.go
│ │     ├── 199_binary-tree-right-side-view.go
│ │     ├── 199_binary-tree-right-side-view_test.go
│ │     ├── 19_remove-nth-node-from-end-of-list.go
│ │     ├── 1_two-sum.go
│ │     ├── 200_number-of-islands.go
│ │     ├── 200_number-of-islands_test.go
│ │     ├── 202_happy-number.go
│ │     ├── 202_happy-number_test.go
│ │     ├── 203_remove-linked-list-elements.go
│ │     ├── 203_remove-linked-list-elements_test.go
│ │     ├── 207_course-schedule.go
│ │     ├── 207_course-schedule_test.go
│ │     ├── 208_implement-trie-prefix-tree.go
│ │     ├── 208_implement-trie-prefix-tree_test.go
│ │     ├── 209_minimum-size-subarray-sum.go
│ │     ├── 209_minimum-size-subarray-sum_test.go
│ │     ├── 20_valid-parentheses.go
│ │     ├── 20_valid-parentheses.py
│ │     ├── 20_valid-parentheses_test.go
│ │     ├── 2160_minimum-sum-of-four-digit-number-after-splitting-digits.go
│ │     ├── 21_merge-two-sorted-lists.go
│ │     ├── 225_implement-stack-using-queues.go
│ │     ├── 226_invert-binary-tree.go
│ │     ├── 22_generate-parentheses.go
│ │     ├── 22_generate-parentheses_test.go
│ │     ├── 230_kth-smallest-element-in-a-bst.go
│ │     ├── 232_implement-queue-using-stacks.go
│ │     ├── 234_palindrome-linked-list.go
│ │     ├── 234_palindrome-linked-list_test.go
│ │     ├── 236_lowest-common-ancestor-of-a-binary-tree.go
│ │     ├── 236_lowest-common-ancestor-of-a-binary-tree_test.go
│ │     ├── 238_product-of-array-except-self.go
│ │     ├── 238_product-of-array-except-self_test.go
│ │     ├── 239_sliding-window-maximum.go
│ │     ├── 239_sliding-window-maximum_test.go
│ │     ├── 23_merge-k-sorted-lists.go
│ │     ├── 240_search-a-2d-matrix-ii.go
│ │     ├── 240_search-a-2d-matrix-ii_test.go
│ │     ├── 242_valid-anagram.go
│ │     ├── 242_valid-anagram_test.go
│ │     ├── 24_swap-nodes-in-pairs.go
│ │     ├── 24_swap-nodes-in-pairs_test.go
│ │     ├── 25_reverse-nodes-in-k-group.go
│ │     ├── 25_reverse-nodes-in-k-group_test.go
│ │     ├── 26_remove-duplicates-from-sorted-array.go
│ │     ├── 26_remove-duplicates-from-sorted-array_test.go
│ │     ├── 2724_sort-by.go
│ │     ├── 27_remove-element.go
│ │     ├── 27_remove-element.py
│ │     ├── 27_remove-element_test.go
│ │     ├── 283_move-zeroes.go
│ │     ├── 283_move-zeroes.py
│ │     ├── 283_move-zeroes_test.go
│ │     ├── 28_find-the-index-of-the-first-occurrence-in-a-string.go
│ │     ├── 28_find-the-index-of-the-first-occurrence-in-a-string_test.go
│ │     ├── 2_add-two-numbers.go
│ │     ├── 2_add-two-numbers_test.go
│ │     ├── 3158_find-the-xor-of-numbers-which-appear-twice.go
│ │     ├── 3158_find-the-xor-of-numbers-which-appear-twice_test.go
│ │     ├── 33_search-in-rotated-sorted-array.go
│ │     ├── 33_search-in-rotated-sorted-array_test.go
│ │     ├── 344_reverse-string.go
│ │     ├── 344_reverse-string_test.go
│ │     ├── 349_intersection-of-two-arrays.go
│ │     ├── 349_intersection-of-two-arrays_test.go
│ │     ├── 35_search-insert-position.go
│ │     ├── 376_wiggle-subsequence.go
│ │     ├── 383_ransom-note.go
│ │     ├── 394_decode-string.go
│ │     ├── 394_decode-string_test.go
│ │     ├── 39_combination-sum.go
│ │     ├── 39_combination-sum_test.go
│ │     ├── 3_longest-substring-without-repeating-characters.go
│ │     ├── 3_longest-substring-without-repeating-characters_test.go
│ │     ├── 416_partition-equal-subset-sum.go
│ │     ├── 416_partition-equal-subset-sum_test.go
│ │     ├── 41_first-missing-positive.go
│ │     ├── 41_first-missing-positive_test.go
│ │     ├── 437_path-sum-iii.go
│ │     ├── 437_path-sum-iii_test.go
│ │     ├── 438_find-all-anagrams-in-a-string.go
│ │     ├── 438_find-all-anagrams-in-a-string_test.go
│ │     ├── 454_4sum-ii.go
│ │     ├── 454_4sum-ii_test.go
│ │     ├── 455_assign-cookies.go
│ │     ├── 455_assign-cookies_test.go
│ │     ├── 46_permutations.go
│ │     ├── 46_permutations_test.go
│ │     ├── 47_permutations-ii.go
│ │     ├── 47_permutations-ii_test.go
│ │     ├── 48_rotate-image.go
│ │     ├── 494_target-sum.go
│ │     ├── 494_target-sum_test.go
│ │     ├── 49_group-anagrams.go
│ │     ├── 509_fibonacci-number.go
│ │     ├── 509_fibonacci-number_test.go
│ │     ├── 51_n-queens.go
│ │     ├── 51_n-queens_test.go
│ │     ├── 53_maximum-subarray.go
│ │     ├── 53_maximum-subarray_test.go
│ │     ├── 541_reverse-string-ii.go
│ │     ├── 541_reverse-string-ii_test.go
│ │     ├── 543_diameter-of-binary-tree.go
│ │     ├── 54_spiral-matrix.go
│ │     ├── 54_spiral-matrix_test.go
│ │     ├── 560_subarray-sum-equals-k.go
│ │     ├── 560_subarray-sum-equals-k_test.go
│ │     ├── 56_merge-intervals.go
│ │     ├── 56_merge-intervals_test.go
│ │     ├── 59_spiral-matrix-ii.go
│ │     ├── 59_spiral-matrix-ii.py
│ │     ├── 59_spiral-matrix-ii_test.go
│ │     ├── 62_unique-paths.go
│ │     ├── 62_unique-paths_test.go
│ │     ├── 63_unique-paths-ii.go
│ │     ├── 63_unique-paths-ii_test.go
│ │     ├── 6_zigzag-conversion.go
│ │     ├── 704_binary-search.go
│ │     ├── 704_binary-search.py
│ │     ├── 704_binary-search_test.go
│ │     ├── 707_design-linked-list.go
│ │     ├── 70_climbing-stairs.go
│ │     ├── 70_climbing-stairs_test.go
│ │     ├── 739_daily-temperatures.go
│ │     ├── 73_set-matrix-zeroes.go
│ │     ├── 746_min-cost-climbing-stairs.go
│ │     ├── 746_min-cost-climbing-stairs_test.go
│ │     ├── 74_search-a-2d-matrix.go
│ │     ├── 74_search-a-2d-matrix_test.go
│ │     ├── 76_minimum-window-substring.go
│ │     ├── 76_minimum-window-substring_test.go
│ │     ├── 78_subsets.go
│ │     ├── 78_subsets_test.go
│ │     ├── 79_word-search.go
│ │     ├── 79_word-search_test.go
│ │     ├── 844_backspace-string-compare.go
│ │     ├── 844_backspace-string-compare.py
│ │     ├── 844_backspace-string-compare_test.go
│ │     ├── 84_largest-rectangle-in-histogram.go
│ │     ├── 84_largest-rectangle-in-histogram_test.go
│ │     ├── 94_binary-tree-inorder-traversal.go
│ │     ├── 94_binary-tree-inorder-traversal_test.go
│ │     ├── 977_squares-of-a-sorted-array.go
│ │     ├── 977_squares-of-a-sorted-array_test.go
│ │     ├── 98_validate-binary-search-tree.go
│ │     ├── 994_rotting-oranges.go
│ │     ├── 994_rotting-oranges_test.go
│ │     ├── LCR 021_SLwz0R.go
│ │     ├── LCR 021_SLwz0R_test.go
│ │     ├── LCR 024_UHnkqh.go
│ │     ├── LCR 024_UHnkqh_test.go
│ │     ├── LCR 036_8Zf90G.go
│ │     ├── LCR 036_8Zf90G_test.go
│ │     ├── LCR 060_g5c51o.go
│ │     ├── LCR 146_shun-shi-zhen-da-yin-ju-zhen-lcof.go
│ │     ├── LCR 146_shun-shi-zhen-da-yin-ju-zhen-lcof_test.go
│ │     ├── define.go
│ │     ├── doc
│ │     │ ├── content
│ │     │ │ ├── 1002_find-common-characters.md
│ │     │ │ ├── 101_symmetric-tree.md
│ │     │ │ ├── 102_binary-tree-level-order-traversal.md
│ │     │ │ ├── 1047_remove-all-adjacent-duplicates-in-string.md
│ │     │ │ ├── 1049_last-stone-weight-ii.md
│ │     │ │ ├── 104_maximum-depth-of-binary-tree.md
│ │     │ │ ├── 105_construct-binary-tree-from-preorder-and-inorder-traversal.md
│ │     │ │ ├── 108_convert-sorted-array-to-binary-search-tree.md
│ │     │ │ ├── 114_flatten-binary-tree-to-linked-list.md
│ │     │ │ ├── 11_container-with-most-water.md
│ │     │ │ ├── 1209_remove-all-adjacent-duplicates-in-string-ii.md
│ │     │ │ ├── 124_binary-tree-maximum-path-sum.md
│ │     │ │ ├── 138_copy-list-with-random-pointer.md
│ │     │ │ ├── 141_linked-list-cycle.md
│ │     │ │ ├── 142_linked-list-cycle-ii.md
│ │     │ │ ├── 144_binary-tree-preorder-traversal.md
│ │     │ │ ├── 146_lru-cache.md
│ │     │ │ ├── 148_sort-list.md
│ │     │ │ ├── 151_reverse-words-in-a-string.md
│ │     │ │ ├── 155_min-stack.md
│ │     │ │ ├── 15_3sum.md
│ │     │ │ ├── 160_intersection-of-two-linked-lists.md
│ │     │ │ ├── 17_letter-combinations-of-a-phone-number.md
│ │     │ │ ├── 189_rotate-array.md
│ │     │ │ ├── 18_4sum.md
│ │     │ │ ├── 199_binary-tree-right-side-view.md
│ │     │ │ ├── 19_remove-nth-node-from-end-of-list.md
│ │     │ │ ├── 1_two-sum.md
│ │     │ │ ├── 200_number-of-islands.md
│ │     │ │ ├── 202_happy-number.md
│ │     │ │ ├── 203_remove-linked-list-elements.md
│ │     │ │ ├── 206_reverse-linked-list.md
│ │     │ │ ├── 207_course-schedule.md
│ │     │ │ ├── 208_implement-trie-prefix-tree.md
│ │     │ │ ├── 209_minimum-size-subarray-sum.md
│ │     │ │ ├── 20_valid-parentheses.md
│ │     │ │ ├── 2160_minimum-sum-of-four-digit-number-after-splitting-digits.md
│ │     │ │ ├── 21_merge-two-sorted-lists.md
│ │     │ │ ├── 225_implement-stack-using-queues.md
│ │     │ │ ├── 226_invert-binary-tree.md
│ │     │ │ ├── 22_generate-parentheses.md
│ │     │ │ ├── 230_kth-smallest-element-in-a-bst.md
│ │     │ │ ├── 232_implement-queue-using-stacks.md
│ │     │ │ ├── 234_palindrome-linked-list.md
│ │     │ │ ├── 236_lowest-common-ancestor-of-a-binary-tree.md
│ │     │ │ ├── 238_product-of-array-except-self.md
│ │     │ │ ├── 239_sliding-window-maximum.md
│ │     │ │ ├── 23_merge-k-sorted-lists.md
│ │     │ │ ├── 240_search-a-2d-matrix-ii.md
│ │     │ │ ├── 242_valid-anagram.md
│ │     │ │ ├── 24_swap-nodes-in-pairs.md
│ │     │ │ ├── 25_reverse-nodes-in-k-group.md
│ │     │ │ ├── 26_remove-duplicates-from-sorted-array.md
│ │     │ │ ├── 2724_sort-by.md
│ │     │ │ ├── 27_remove-element.md
│ │     │ │ ├── 283_move-zeroes.md
│ │     │ │ ├── 28_find-the-index-of-the-first-occurrence-in-a-string.md
│ │     │ │ ├── 2_add-two-numbers.md
│ │     │ │ ├── 3158_find-the-xor-of-numbers-which-appear-twice.md
│ │     │ │ ├── 33_search-in-rotated-sorted-array.md
│ │     │ │ ├── 344_reverse-string.md
│ │     │ │ ├── 349_intersection-of-two-arrays.md
│ │     │ │ ├── 35_search-insert-position.md
│ │     │ │ ├── 376_wiggle-subsequence.md
│ │     │ │ ├── 383_ransom-note.md
│ │     │ │ ├── 394_decode-string.md
│ │     │ │ ├── 39_combination-sum.md
│ │     │ │ ├── 3_longest-substring-without-repeating-characters.md
│ │     │ │ ├── 416_partition-equal-subset-sum.md
│ │     │ │ ├── 41_first-missing-positive.md
│ │     │ │ ├── 437_path-sum-iii.md
│ │     │ │ ├── 438_find-all-anagrams-in-a-string.md
│ │     │ │ ├── 454_4sum-ii.md
│ │     │ │ ├── 455_assign-cookies.md
│ │     │ │ ├── 46_permutations.md
│ │     │ │ ├── 47_permutations-ii.md
│ │     │ │ ├── 48_rotate-image.md
│ │     │ │ ├── 494_target-sum.md
│ │     │ │ ├── 509_fibonacci-number.md
│ │     │ │ ├── 51_n-queens.md
│ │     │ │ ├── 53_maximum-subarray.md
│ │     │ │ ├── 541_reverse-string-ii.md
│ │     │ │ ├── 543_diameter-of-binary-tree.md
│ │     │ │ ├── 54_spiral-matrix.md
│ │     │ │ ├── 560_subarray-sum-equals-k.md
│ │     │ │ ├── 56_merge-intervals.md
│ │     │ │ ├── 59_spiral-matrix-ii.md
│ │     │ │ ├── 62_unique-paths.md
│ │     │ │ ├── 63_unique-paths-ii.md
│ │     │ │ ├── 6_zigzag-conversion.md
│ │     │ │ ├── 704_binary-search.md
│ │     │ │ ├── 707_design-linked-list.md
│ │     │ │ ├── 70_climbing-stairs.md
│ │     │ │ ├── 739_daily-temperatures.md
│ │     │ │ ├── 73_set-matrix-zeroes.md
│ │     │ │ ├── 746_min-cost-climbing-stairs.md
│ │     │ │ ├── 74_search-a-2d-matrix.md
│ │     │ │ ├── 76_minimum-window-substring.md
│ │     │ │ ├── 78_subsets.md
│ │     │ │ ├── 844_backspace-string-compare.md
│ │     │ │ ├── 84_largest-rectangle-in-histogram.md
│ │     │ │ ├── 94_binary-tree-inorder-traversal.md
│ │     │ │ ├── 977_squares-of-a-sorted-array.md
│ │     │ │ ├── 98_validate-binary-search-tree.md
│ │     │ │ ├── 994_rotting-oranges.md
│ │     │ │ ├── LCP 55_PTXy4P.md
│ │     │ │ ├── LCR 021_SLwz0R.md
│ │     │ │ ├── LCR 024_UHnkqh.md
│ │     │ │ ├── LCR 032_dKk3P7.md
│ │     │ │ ├── LCR 036_8Zf90G.md
│ │     │ │ ├── LCR 060_g5c51o.md
│ │     │ │ ├── LCR 078_vvXgSW.md
│ │     │ │ ├── LCR 088_GzCJIP.md
│ │     │ │ ├── LCR 107_2bCMpM.md
│ │     │ │ ├── LCR 146_shun-shi-zhen-da-yin-ju-zhen-lcof.md
│ │     │ │ ├── [121]best-time-to-buy-and-sell-stock.md
│ │     │ │ ├── [144]二叉树的前序遍历.md
│ │     │ │ ├── [1]two-sum.md
│ │     │ │ ├── [206]reverse-linked-list.md
│ │     │ │ ├── [21]merge-two-sorted-lists.md
│ │     │ │ └── [704]binary-search.md
│ │     │ ├── note
│ │     │ │ ├── 141_linked-list-cycle.md
│ │     │ │ ├── 509_fibonacci-number.md
│ │     │ │ ├── 70_climbing-stairs.md
│ │     │ │ ├── 746_min-cost-climbing-stairs.md
│ │     │ │ └── 977_squares-of-a-sorted-array.md
│ │     │ └── solution
│ │     │     ├── bao-mu-shi-ti-jie-shou-ba-shou-da-tong-tuo-bu-pai-.lcv
│ │     │     ├── bi-jiao-han-tui-ge-de-zi-fu-chuan-by-leetcode-solu.lcv
│ │     │     ├── chang-du-zui-xiao-de-zi-shu-zu-by-leetcode-solutio.lcv
│ │     │     ├── chu-zi-shen-yi-wai-shu-zu-de-cheng-ji-by-leetcode-.lcv
│ │     │     ├── cong-qian-xu-yu-zhong-xu-bian-li-xu-lie-gou-zao-9.lcv
│ │     │     ├── dao-yu-shu-liang-by-leetcode.lcv
│ │     │     ├── die-dai-fa-by-jason-2.lcv
│ │     │     ├── dui-cheng-er-cha-shu-by-leetcode-solution.lcv
│ │     │     ├── er-cha-shu-de-ceng-xu-bian-li-by-leetcode-solution.lcv
│ │     │     ├── er-cha-shu-de-you-shi-tu-by-leetcode-solution.lcv
│ │     │     ├── er-cha-shu-de-zhi-jing-by-leetcode-solution.lcv
│ │     │     ├── er-cha-shu-de-zhong-xu-bian-li-by-leetcode-solutio.lcv
│ │     │     ├── er-cha-shu-de-zui-da-shen-du-by-leetcode-solution.lcv
│ │     │     ├── er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetc-2.lcv
│ │     │     ├── er-cha-shu-zhan-kai-wei-lian-biao-by-leetcode-solu.lcv
│ │     │     ├── er-cha-shu-zhong-de-zui-da-lu-jing-he-by-ikaruga.lcv
│ │     │     ├── er-cha-shu-zhong-de-zui-da-lu-jing-he-by-leetcode-.lcv
│ │     │     ├── er-cha-sou-suo-shu-zhong-di-kxiao-de-yua-8o07.lcv
│ │     │     ├── fei-bo-na-qi-shu-by-leetcode-solution-o4ze.lcv
│ │     │     ├── fu-lan-de-ju-zi-by-leetcode-solution.lcv
│ │     │     ├── golanghui-su-by-sealyun.lcv
│ │     │     ├── gua-hao-sheng-cheng-by-leetcode-solution.lcv
│ │     │     ├── he-bing-kge-pai-xu-lian-biao-by-leetcode-solutio-2.lcv
│ │     │     ├── he-bing-qu-jian-by-leetcode-solution.lcv
│ │     │     ├── he-wei-kde-zi-shu-zu-by-leetcode-solution.lcv
│ │     │     ├── huan-xing-lian-biao-de-jie-dian-ding-wei-a68a.lcv
│ │     │     ├── huan-xing-lian-biao-ii-by-leetcode-solution.lcv
│ │     │     ├── hui-su-bu-hui-xie-tao-lu-zai-ci-pythonja-wcdw.lcv
│ │     │     ├── hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-2.lcv
│ │     │     ├── hui-su-tao-lu-miao-sha-nhuang-hou-shi-pi-mljv.lcv
│ │     │     ├── hui-wen-lian-biao-by-leetcode-solution.lcv
│ │     │     ├── jiang-you-xu-shu-zu-zhuan-huan-wei-er-cha-sou-s-33.lcv
│ │     │     ├── jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-s1cx.lcv
│ │     │     ├── ju-zhen-zhi-ling-by-leetcode-solution-9ll7.lcv
│ │     │     ├── k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio.lcv
│ │     │     ├── kge-yi-zu-fan-zhuan-lian-biao-by-powcai.lcv
│ │     │     ├── liang-chong-fang-fa-cong-o52mn-dao-omnfu-3ezz.lcv
│ │     │     ├── linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-.lcv
│ │     │     ├── lu-jing-zong-he-iii-by-leetcode-solution-z9td.lcv
│ │     │     ├── mei-ri-wen-du-by-leetcode-solution.lcv
│ │     │     ├── mei-xiang-ming-bai-yi-ge-shi-pin-jiang-t-nvsq.lcv
│ │     │     ├── mian-shi-ti-29-shun-shi-zhen-da-yin-ju-zhen-she-di.lcv
│ │     │     ├── nhuang-hou-by-leetcode-solution.lcv
│ │     │     ├── pai-xu-lian-biao-by-leetcode-solution.lcv
│ │     │     ├── qian-zhui-he-ha-xi-biao-cong-liang-ci-bi-4mwr.lcv
│ │     │     ├── quan-pai-lie-by-leetcode-solution-2.lcv
│ │     │     ├── que-shi-de-di-yi-ge-zheng-shu-by-leetcode-solution.lcv
│ │     │     ├── shan-chu-lian-biao-de-dao-shu-di-n-ge-ji-ydte.lcv
│ │     │     ├── shan-chu-zi-fu-chuan-zhong-de-suo-you-xi-4ohr.lcv
│ │     │     ├── she-ji-lian-biao-by-leetcode-solution-abix.lcv
│ │     │     ├── shi-xian-strstr-by-leetcode-solution-ds6y.lcv
│ │     │     ├── shi-xian-trie-qian-zhui-shu-by-leetcode-ti500.lcv
│ │     │     ├── shua-chuan-lc-shuang-bai-po-su-jie-fa-km-tb86.lcv
│ │     │     ├── shun-shi-zhen-da-yin-ju-zhen-by-leetcode-solution.lcv
│ │     │     ├── sou-suo-cha-ru-wei-zhi-by-leetcode-solution.lcv
│ │     │     ├── sou-suo-er-wei-ju-zhen-ii-by-leetcode-so-9hcx.lcv
│ │     │     ├── tu-jie-guan-fang-tui-jian-ti-jie-shan-ch-x8iz.lcv
│ │     │     ├── tu-jie-kge-yi-zu-fan-zhuan-lian-biao-by-user7208t.lcv
│ │     │     ├── xuan-zhuan-tu-xiang-by-leetcode-solution-vu3m.lcv
│ │     │     ├── yan-zheng-er-cha-sou-suo-shu-by-leetcode-solution.lcv
│ │     │     ├── you-xu-shu-zu-de-ping-fang-by-leetcode-solution.lcv
│ │     │     ├── zi-fu-chuan-jie-ma-by-leetcode-solution.lcv
│ │     │     ├── zu-he-zong-he-by-leetcode-solution.lcv
│ │     │     ├── zui-da-zi-xu-he-by-leetcode-solution.lcv
│ │     │     ├── zui-xiao-fu-gai-zi-chuan-by-leetcode-solution.lcv
│ │     │     └── zui-xiao-zhan-by-leetcode-solution.lcv
│ │     └── test_framework.go
│ ├── list
│ │ ├── deleteDuplicates.go
│ │ ├── deleteDuplicates_test.go
│ │ ├── list_node.go
│ │ ├── reversePrint.go
│ │ └── reversePrint_test.go
│ ├── list_node
│ │ ├── list_node.go
│ │ └── list_node_test.go
│ ├── listnode.go
│ ├── monotonic_queue
│ │ └── monotonic_queue.go
│ ├── recursion
│ │ ├── isPowe.go
│ │ └── isPowe_test.go
│ ├── stack.go
│ ├── string
│ │ ├── addBinary.go
│ │ ├── addBinary_test.go
│ │ ├── canConstruct.go
│ │ ├── canConstruct_test.go
│ │ ├── findTheDifference.go
│ │ ├── findTheDifference_test.go
│ │ ├── firstUniqChar.go
│ │ ├── firstUniqChar_test.go
│ │ ├── lengthOfLastWord.go
│ │ ├── lengthOfLastWord_test.go
│ │ ├── longestCommonPrefix.go
│ │ ├── longestCommonPrefix_test.go
│ │ ├── reverseVowels.go
│ │ ├── reverseVowels_test.go
│ │ ├── romanToInt.go
│ │ ├── strStr.go
│ │ ├── strStr_test.go
│ │ └── string_test.go
│ └── tree
│     ├── hasPathSum.go
│     ├── isSameTree.go
│     ├── isSubtree.go
│     ├── isSymmetric.go
│     ├── maxDepth.go
│     ├── maxPathSum.go
│     ├── minDepth.go
│     ├── treenode.go
│     └── treenode_test.go
├── main.go
└── testutil
    ├── int.go
    ├── slice.go
    └── slice_test.go

```

ref: https://github.com/youngyangyang04/leetcode-master?tab=readme-ov-file


## 常用操作

### 取数值各个位上的单数

```go
func IterDigest(n int, do func(digest int)) {
	for n != 0 {
		do(n % 10)
		n /= 10
	}
}
```

例如,取数值各个位上的单数之和:
```go
func getSum(n int) (sum int) {
	iterDigest(n, func(digest int) {
		sum += digest * digest
	})
	return
}
```

