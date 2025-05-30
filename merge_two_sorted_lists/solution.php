<?php

namespace MergeTwoSortedLists;



 class ListNode {
     public $val = 0;
     public $next = null;
     function __construct($val = 0, $next = null) {
         $this->val = $val;
         $this->next = $next;
     }
 }

class Solution {
    /**
     * @param ListNode $list1
     * @param ListNode $list2
     * @return ListNode
     */
    function mergeTwoLists(ListNode $list1, ListNode $list2): ListNode
    {
        if ($list1->val === null) {
            return $list2;
        }

        if ($list2->val === null) {
            return $list1;
        }

        if ($list1->val < $list2->val) {
            $list1->next = $this->mergeTwoLists($list1->next, $list2);
            return $list1;
        }

        $list2->next = $this->mergeTwoLists($list1, $list2->next);
        return $list2;
    }
}
