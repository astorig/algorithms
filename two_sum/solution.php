<?php
namespace TwoSum;

class Solution {
    /**
     * @param array $nums
     * @param int $target
     * @return array
     */
    public function twoSum (array $nums, int $target): array
    {
        $result = [];
        $array = [];
        foreach ($nums as $key=>$num) {
            $array[$target - $num] = $key;
            if (isset($array[$num])) {
                $result = [$array[$num], $key];
                break;
            }
        }

        return $result;
    }
}
$exc = new Solution();

$nums = [2, 2, 3, 4, 5];
$target = 5;


$result = $exc->twoSum($nums, $target);

var_dump($result);