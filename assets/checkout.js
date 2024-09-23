const stripe = Stripe(
  "pk_test_51PtsLPP4MxIMgAth8ymcxGIUDZTilzzf9nOFiwkmXTKyT149RsxH4kXW9CKvUEt6jI02Pq5h8kVtaXfNlMI4RRXF00bT3of6po",
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
    confirmParams: { return_url: "http://localhost:7331/settings/billing" },
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
