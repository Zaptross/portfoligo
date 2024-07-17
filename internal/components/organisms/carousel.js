function initCarousel(carouselInner, carouselButtons) {
  const carouselItems = carouselInner.getElementsByClassName("carousel-item");
  const totalSlides = carouselItems.length;
  carouselInner.dataset.index = 0;

  function showSlide(index) {
    const currentIndex = index % totalSlides;
    const translateValue = currentIndex * 100 + "%";
    carouselInner.style.transform = "translateX(-" + translateValue + ")";
  }

  function prevSlide() {
    showSlide(nextIndex(-1));
  }

  function nextSlide() {
    showSlide(nextIndex(1));
  }

  function nextIndex(direction) {
    const currentIndex = parseInt(carouselInner.dataset.index);
    let nextIndex = (currentIndex + direction) % totalSlides;

    if (nextIndex < 0) {
      nextIndex = totalSlides - 1;
    }

    carouselInner.dataset.index = nextIndex;
    return nextIndex;
  }

  // Attach click event listeners to buttons
  carouselButtons[0].addEventListener("click", prevSlide);
  carouselButtons[1].addEventListener("click", nextSlide);

  return {
    showSlide,
    prevSlide,
    nextSlide,
  };
}
