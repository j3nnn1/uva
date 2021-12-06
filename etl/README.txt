
- Descargar archivos desde la url del README principal
- Ordernar los archivos de menor a mayor fecha. (separar Menu/Data/Text to columns, rellenar con 00 (=TEXT(C2798, "00")), concatenar las fechas  y ordenar por fechas = sort Conscotiza.csv | uniq > Conscotiza_menor_mayor_fecha_usd.csv)
- Eliminar header del archivo de cotizacion del usd, borrar filas anteriores al 31 de marzo del 2016 y guardarlo en formato csv.

ConscotizaWITHOUTHEADER.csv
cer-uva-uvi-diarios.csv


ConscotizaWITHOUTHEADER.csv - columnas: 20160331,31-03-2016,31,03,2016,14.539,14.889 
filas a partir de: 31-03-2016 ( 31 de marzo del 2016 se creo la UVA.)
output: filledDates.csv


cer-uva-uvi-diarios.csv - columnas: indice_tiempo,cer_diario,uva_diario,uvi_diario,usd_date_cotizacion,usd_date_cotizacion_int,usd_compra,usd_venta
filas a partir de: 31-03-2016 ( 31 de marzo del 2016 se creo la UVA.)
nombre del archivo unificado.

- Los datos de cierre de cotizacion del dolar tienen gap por los fines de semana, en ese caso se deja el valor del dolar del viernes o del lunes. para ello se uso un script de go para llenar los gaps.

```
go build makeConsecutiveDays.go
go run 00makeConsecutiveDays.go  > logs.txt
```

Esto genera este archivo: "filledDates.csv" copiar las columnas fecha, usd de venta y unificar con la cotizacion uva en un nuevo archivo: "uva_uvi_usd_final.csv".
La columna 0 corresponde a la fecha que establece el orden de tiempo.

2016-03-31,5.5636,14.05,14.05,20160331,14.539,14.889,19.3557

@TODO agregar variable para definir columnas a analizar en busqueda de un gap, actualmente esta en un numero constante.

-  Con el archivo anterior generado se utiliza para transformar el json que necesita para hacer el grafico  de lineas.
   Generando un archivo nuevo con el formato json del template de D3.js: graphLineTemplate.json

```
go build makeConsecutiveDays.go
./binario > logs.txt

OR

go run 00makeConsecutiveDays.go  > logs.txt
```

   Para cambiar que columna tomar del csv se modifica en la definición de estas variables: 
   Es importante la key del map deben ser iguales, y tener sus correspondencias.
   La columna 0 corresponde a la fecha que establece el orden de tiempo.

```
    header := map[string]Line {
        "USD":{},
        "ARG":{},
    }
    body := map[string][]Point {
        "USD": {},
        "ARG": {},
    }
    variables := map[string]int {
        "USD": 6, // numero de columna teniendo que la primera columna es 0
        "ARG": 2, // numero de columna teniendo que la primera columna es 0
    }
```

- Ojito el Json template generado tiene mayusculas en los keys, llevar todo a minuscula, "series" "value" y "name".

- para probar ng serve

- para generar distribucion y guardar en gh-pages

ng build --prod --base-href "https://j3nnn1.github.io/uva/"

- copiar archivos de la carpeta ./dist a la carpeta raiz en el branch gh-pages. aprox 5, 6 files 

- push origin gh-pages


- nvm
- node VERSION
- angular VERSION 7.2.15

npm i @angular/core@7.2.15