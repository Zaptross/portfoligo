function configureSearch(collectionParent, dataProperty, searchInput) {
  window.addEventListener("load", function () {
    if (location.search) {
      const search = decodeURI(location.search).replace("?q=", "");
      searchInput.value = search;
      onSearch({target: {value: search}});
    }
  });

  function onSearch(event) {
    const searchValue = event.target.value;
    const collection = Array.from(collectionParent?.children);

    collection.forEach(function (element) {
      const data = element.dataset[dataProperty].toLowerCase();
      const elementShouldBeHidden = !data.includes(searchValue.toLowerCase());
      element.classList.toggle("hidden", elementShouldBeHidden);
    });

    // After the search, check if we should show the back to top button
    checkShowBackToTop();
  }

  searchInput?.addEventListener("keyup", onSearch);
}
