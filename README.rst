ITRI-TTS client
===============

Install
-------

.. code:: shell

    go get -u github.com/myyang/tts

Usage
-----

Example:

.. code:: go

    package main

    import (
        "fmt"
        "github.com/myyang/tts"
    )

    func main() {
        s := "工研院的中英交錯Sentence做得比Google好喔，詳見測試檔"
        url := tts.ConvertSimple("account", "password", s)
        fmt.Printf("%s", url)
        // http.Get(url)
    }

Example output:

.. code:: shell

   http://tts.itri.org.tw/TTSservice/download/account/146355769925921.wav
