<?php
namespace InvertBinaryTree;

class TreeNode {
    public $val;
    public $left;
    public $right;

    public function __construct($val = 0, $left = null, $right = null)
    {

        $this->right = $right;
        $this->left = $left;
        $this->val = $val;
    }
}

class Solution {

    /**
     * @param TreeNode $root
     * @return TreeNode
     */
    public function invertTree (TreeNode $root): TreeNode
    {
        if ($root->left === null && $root->right === null) {
            return $root;
        }
        $left = $root->left;
        $root->left = $root->right;
        $root->right = $left;

        $this->invertTree($root->left);
        $this->invertTree($root->right);

        return $root;
    }
}
