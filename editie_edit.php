<?php
if ($_SERVER["REQUEST_METHOD"] === "POST") {

    $data_editie = $_POST["data_editie"];
    $numar_editie = $_POST["numar_editie"];

    echo "Data Editie: " . $data_editie . "\n";
    echo "Numar Editie: " . $numar_editie . "\n";

    
}
?>