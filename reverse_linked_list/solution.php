<?php
namespace ReverseLinkedList;

class ListNode {
    public $val = 0;
    public $next = null;

    public function __construct($val, $next)
    {

        $this->val = $val;
        $this->next = $next;
    }
}
class Solution {
    /**
     * @param ListNode $head
     * @return ListNode
     */
    public function reverseList (ListNode $head): ListNode
    {
        $prev = null;

        while ($head !== null) {
            $next = $head->next;
            $head->next = $prev;
            $prev = $head;
            $head = $next;
        }

        return $prev;
    }
}
