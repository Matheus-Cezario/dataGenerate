# Gerador de dados aleatórios em GO
WebAssembly escrito em GO que gera dados aleatórios a partir de um JSON de configuração.

## Configuração do WebAssembly
```html
<html>  
    <head>
        <meta charset="utf-8"/>
        <script src="wasm_exec.js"></script>

    </head>
    <body>

    </body>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });

        function Generate(data,qtd){;
            result = GOGenerate(JSON.stringify(data),qtd);
            return JSON.parse(result)
        }

    </script>
</html>
```

## JSON de configuração

```json
{
  "Nome do atributo": "Nome da função geradora:parametros separdos por ':'"
}
```
### Exemplo
```json
{
  "nome":"name:m",
  "id":"number:10:25",
  "cidade":"choise:Ouro Branco:Ouro Preto:Belo Horizonte",
  "produtos":"choiseAny:r:Bala:Sorvete:Ovo:Arroz:Batata:Chocolate"
}
```

![image](https://user-images.githubusercontent.com/44527383/160287722-afcf1f4f-4b4f-4624-9b17-6022f8c1962c.png)
