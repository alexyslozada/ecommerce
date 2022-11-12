# Proceso de Webhook en el backend

Para poder procesar correctamente los pagos de PayPal, la mejor práctica es procesar los eventos que envía PayPal hacia nuestro servidor. Aquí aprenderemos a realizarlo de forma correcta.

## Qué es un webhook?

Los webhook son eventos que desencadenan una acción. Generalmente, se confunden con las API REST, ya que funcionan bajo la misma arquitectura, sin embargo, la diferencia radica en que los webhook se disparan desde el sitio que quieres consultar y te notifican cuando se ha realizado alguna acción. Mientras que las API REST tú debes ejecutar la petición para saber si algo ha cambiado.

Por lo tanto, lo que tú haces es crear un endpoint que recibirá las notificaciones que envíe el otro servicio para procesarlas y así no debes estar haciendo solicitudes para saber si algo ocurrió en el servicio que estás consumiendo.

## 1. Activando los webhook en PayPal

Vamos a activar los webhook en PayPal para que nos notifique cada vez que se ha procesado un pago.

- Vamos al dashboard en: [developer.paypal.com](http://developer.paypal.com)
- Seleccionamos en Dashboard `My apps & credentials`
- Seleccionamos nuestra app (si no has creado ninguna, selecciona la que viene por default)
- Vamos a la sección de `Sandbox Webhooks`
- Damos clic en la opción `Add webhook`
- Colocaremos la URL completa de nuestro webhook (si aún no lo tienes no te preocupes, lo podremos modificar luego). Nota: debe ser https.
- Aunque podemos seleccionar todos los webhook, por ahora solo usaremos uno:
    - `Payment Capture Completed` (este nos sirve para saber cuando se ha pagado un producto)
- Una vez creado nos muestra el webhook con su respectivo ID.

## 2. Webhook simulator

PayPal nos ofrece una forma sencilla de hacer pruebas de nuestro webhook sin tener que realizar todo el proceso de compra o suscripción. Y nos sirve para poder validar si estamos recibiendo los datos para procesarlos. Para utilizarlo debemos:

- Vamos al dashboard en: [developer.paypal.com](http://developer.paypal.com)
- En la sección `Mock` seleccionamos `Webhooks simulator`
- Digitamos la URL de nuestro webhook (debe ser https).
- Seleccionamos el webhook que deseamos recibir.
- Y damos clic a la opción `Send test`

Si aún no tienes tu endpoint listo para escuchar, no te preocupes, el simulador te muestra cuál fue el contenido del webhook que estás enviando. Por esta razón podrás copiar y pegar ese contenido para hacer un mock en tu código.

## 3. Flujo del proceso del Webhook

[Clic aquí para ver el diagrama](https://miro.com/app/board/uXjVOSPreHY=/?invite_link_id=365415215848)

El proceso está dividido entre lo que ve el cliente y lo que se procesa en el servidor.

### 3.1 Flujo del cliente

El cliente selecciona la opción de pagar con PayPal y una vez ha realizado el pago se redirige a una página informando el proceso exitoso o el proceso de verificación del pago.

### 3.2 Flujo del servidor

El flujo del servidor consiste en los siguientes 4 pasos:

1. PayPal envía una notificación via Webhook a nuestro servidor.
2. Nuestro servidor realiza una petición a PayPal solicitando la verificación de lo que se recibió en el Webhook. Este paso es muy importante, ya que nuestro endpoint es público y cualquier persona podría simular una petición de `Proceso pagado correctamente` , así que el atacante podría simular el pago de un producto y si nosotros no verificamos esa información con PayPal podríamos estar facturando un producto que no han pagado. Más información en [https://developer.paypal.com/api/rest/webhooks/#link-messagesignature](https://developer.paypal.com/api/rest/webhooks/#link-messagesignature)
3. PayPal responde con un OK o un Fallido.
4. Si todo está OK podemos proceder a la facturación.
