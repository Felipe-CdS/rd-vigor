const stripe = Stripe(
  "pk_test_51QFd8CEO8E0VEaQp05Z8YVdz8jLpOjZTMAyJkqMV5Ff7K6gWvKjc3GXBq9GuYOIwY4Gt6DXv4VYRyUFLiBzr7Mbf00wr244bLd",
  { locale: "pt-BR" },
);

window.initializeStripe = async function () {
  const response = await fetch("/create-subscription", { method: "POST" });
  const { SubscriptionID, clientSecret } = await response.json();

  var options = {
    clientSecret,
    appearance: { theme: "stripe" },
  };

  const elements = stripe.elements(options);

  const paymentElement = elements.create("payment");
  paymentElement.mount("#payment-element");

  document
    .querySelector("#payment-form")
    .addEventListener("submit", function (e) {
      handleSubmit(e, elements);
    });

  return "ready";
};

async function handleSubmit(e, elements) {
  e.preventDefault();
  setLoading(true);

  const { error } = await stripe.confirmPayment({
    elements,
    confirmParams: { return_url: `${window.location.origin}/settings/billing` },
  });

  setLoading(false);
}

function setLoading(flag) {
  if (flag) {
    document.querySelector("#submit").innerHTML = "Carregando...";
    document.querySelector("#submit").style.opacity = "0.5";
    document.querySelector("#submit").disabled = true;
    return;
  }
  document.querySelector("#submit").removeAttribute("disabled");
  document.querySelector("#submit").style.opacity = "1";
  document.querySelector("#submit").innerHTML = "Concluir Compra";
}

// ------- Call After transaction -------

window.getStatus = async function (clientSecret) {
  const { paymentIntent } = await stripe.retrievePaymentIntent(clientSecret);

  switch (paymentIntent.status) {
    case "succeeded":
      return 1;
    case "processing":
      return 2;
    case "requires_payment_method":
      return 3;
    default: // Something went wrong.
      return 4;
  }
};
