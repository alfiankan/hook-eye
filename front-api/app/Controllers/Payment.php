<?php

namespace App\Controllers;

use App\Controllers\BaseController;
use App\Models\Payment as ModelsPayment;
use Kafka;

class Payment extends BaseController
{
    protected $paymentModel;
    public function __construct() {
        $this->paymentModel = new ModelsPayment();
    }

    public function makePayment()
    {
        /**
         * event :
         * PAYMENT.CAPTURE
         * PAYMENT.PROCESSED
         * PAYMENT.COMPLETED
         */

        $reqBody = json_decode($this->request->getBody());


        return $this->response->setJSON([
            'success' => true,
            'message' => 'PAYMENT.CAPTURE',
            'data' => [
                'reference_id' => $reqBody->reference_id,
                'note' => $this->paymentModel->processPayment($reqBody),
            ]
        ]);
    }

}
