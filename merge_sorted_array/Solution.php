<?php

namespace merge_sorted_array;

class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer $m
     * @param Integer[] $nums2
     * @param Integer $n
     */
    function merge(&$nums1, int $m, array $nums2, int $n)
    {
        $numOneTarget = $m - 1;
        $numTwoTarget = $n - 1;
        $numOnelastElementKey = count((array) $num1) - 1;

        while ($numTwoTarget >= 0) {
            if ($num1[$numOneTarget] > $num2[$numTwoTarget]) {
                $num1[$numOnelastElementKey] = $num1[$numOneTarget];
                $numOneTarget--;
            } else {
                $num1[$numOnelastElementKey] = $num2[$numTwoTarget];
                $numTwoTarget--;
            }

            $numOnelastElementKey--;
        }
    }
}

$solution = new Solution();
$num1 = [1,2,3,0,0,0];
$m = 3;
$num2 = [2,5,6];
$n = 3;

var_dump($solution->merge($num1, $m, $num2, $n));