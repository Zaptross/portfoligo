function configureSearch(collectionParent, dataProperty, searchInput) {
  searchInput?.addEventListener("keyup", function (event) {
    const searchValue = event.target.value;
    const collection = Array.from(collectionParent?.children);

    collection.forEach(function (element) {
      const data = element.dataset[dataProperty].toLowerCase();
      const elementShouldBeHidden = !data.includes(searchValue.toLowerCase());
      element.classList.toggle("hidden", elementShouldBeHidden);
    });
  });
}
