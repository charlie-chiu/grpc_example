<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greeting.proto

namespace GPBMetadata;

class Greeting
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(hex2bin(
            "0ab7010a0e6772656574696e672e70726f746f12086772656574696e67221c0a0c48656c6c6f52657175657374120c0a046e616d65180120012809222e0a0a48656c6c6f5265706c79120f0a076d657373616765180120012809120f0a076772656574657218022001280932450a0747726565746572123a0a0853617948656c6c6f12162e6772656574696e672e48656c6c6f526571756573741a142e6772656574696e672e48656c6c6f5265706c792200620670726f746f33"
        ), true);

        static::$is_initialized = true;
    }
}
