const asleep = (delay) => {
  return new Promise((resolve) => setTimeout(resolve, delay));
};

(async () => {
  while (!document.getElementsByClassName("pdp-product-title")[0]) {
    await asleep(100);
  }

  const cat = document
    .getElementsByClassName("breadcrumb_item")
    [
      document.getElementsByClassName("breadcrumb_item").length - 2
    ].textContent.trim();
  const brand = document.getElementsByClassName(
    "pdp-link pdp-link_size_s pdp-link_theme_blue pdp-product-brand__brand-link"
  )[0].innerText;
  const name = document
    .getElementsByClassName("pdp-product-title")[0]
    .textContent.trim();

  const priceElem = document.getElementsByClassName("pdp-product-price")[0];

  const price = parseInt(
    priceElem.textContent.replace("Rs.", "").replaceAll(",", "").trim()
  );

  const url = new URL(`https://stormy-refuge-75823.herokuapp.com/match`);
  const params = {
    title: name,
    price: price,
    category: cat,
    brand: brand === "No Brand" ? "" : brand,
  };
  Object.keys(params).forEach((key) =>
    url.searchParams.append(key, params[key])
  );

  let result;

  try {
    const resp = await fetch(url);
    result = await resp.json();
  } catch (err) {
    console.log("Error!");
    return;
  }

  const domElem = document.getElementById("module_product_price_1");
  const newContainer = document.createElement("ul");

  const noImageLink =
    "https://hamrobazaar.obs.ap-southeast-3.myhuaweicloud.com/Assets/NoImage.png";

  for (const item of result) {
    let imageLink = item.imageUrl ? item.imageUrl : noImageLink;

    const html = `
        <div style="padding: 5px;">
            <a href="${item.url}" target="_blank">
                <div style="display: flex;padding: 5px;">
                    <div style="width: 50px">
                        <img style="width:50px;" src="${imageLink}"/>
                    </div>
                    <div style="display: flex; flex-direction: column; padding: 0px 5px">
                        <div style="flex-grow: 1; font-size: 1.5em;">${item.name}</div>
                        <div style="color: #f57224; font-size: 2em;">Rs. ${item.price}</div>
                    </div>
                </div>
            </a>
        </div>
        `;

    const tempContainer = document.createElement("li");

    tempContainer.innerHTML = html;

    newContainer.append(tempContainer);
  }

  newContainer.style["border"] = "1px solid black";

  domElem.append(newContainer);
})();
