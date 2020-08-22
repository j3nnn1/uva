
- Descargar archivos desde la url del README principal
- Ordernar los archivos de menor a mayor fecha.
- Eliminar header del archivo de cotizacion del usd y guardarlo en formato csv.

ConscotizaWITHOUTHEADER.csv
cer-uva-uvi-diarios.csv

- Los datos de cierre de cotizacion del dolar tienen gap por los fines de semana, en ese caso se deja el valor del dolar del viernes o del lunes. para ello se uso un script de go para llenar los gaps.

```
go build makeConsecutiveDays.go
go run 00makeConsecutiveDays.go  > logs.txt
```

Esto genera este archivo: "filledDates.csv" copiar las columnas fecha, usd de venta y unificar con la cotizacion uva en un nuevo archivo: "uva_uvi_usd_final.csv".
La columna 0 corresponde a la fecha que establece el orden de tiempo.
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
