<?php

namespace App\Http\Controllers;

use App\Helpers\Helper;
use Illuminate\Http\Request;

class MetaController extends Controller
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct()
    {
        //
    }
    public function getDocs() {
        return view('docs');
    }

    public function getMeta(Request $request) {
        return Helper::getMetaData($request->link, $request->all);
    }
}
