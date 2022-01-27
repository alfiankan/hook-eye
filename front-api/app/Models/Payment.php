<?php

namespace App\Models;

use CodeIgniter\Model;
use Kafka;

class Payment extends Model
{
   public function processPayment($data): bool
   {
       // send event PAYMENT.PROCESS
       
       $kafka = new Kafka();

       $kafka->send('PAYMENT.CAPTURE', json_encode([
           'event' => 'PAYMENT.CAPTURE',
           'reference_id' => $data->reference_id,
           'total' => $data->total
       ]));
       
       return true;
   }
}
