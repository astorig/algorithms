<?php

namespace Lru_cache;

class LRUCache {
    private const UNFOUND_KEY_VALUE = -1;
    private array $cache = [];
    private int $capacity;

    /**
     * @param Integer $capacity
     */
    public function __construct(int $capacity)
    {
        $this->capacity = $capacity;
    }

    /**
     * @param Integer $key
     * @return Integer
     */
    public function get(int $key): int
    {
        if (array_key_exists($key, $this->cache)) {
            $this->updateKey($key, $this->cache[$key]);

            return $this->cache[$key];
        }

        return self::UNFOUND_KEY_VALUE;
    }

    /**
     * @param Integer $key
     * @param Integer $value
     * @return void
     */
    public function put(int $key, int $value): void
    {
        if (array_key_exists($key, $this->cache)) {
            $this->updateKey($key, $value);
        } else {
            if (count($this->cache) === $this->capacity) {
                $leastUsedKey = array_key_first($this->cache);
                unset($this->cache[$leastUsedKey]);
            }

            $this->cache[$key] = $value;
        }
    }

    /**
     * @param int $key
     * @param int $value
     * @return void
     */
    private function updateKey(int $key, int $value): void
    {
        unset($this->cache[$key]);
        $this->cache[$key] = $value;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * $obj = LRUCache($capacity);
 * $ret_1 = $obj->get($key);
 * $obj->put($key, $value);
 */