window.addEventListener("load", () => {
  const headings = Array.from(document.querySelectorAll("a[name]")).filter(
    // If the anchor has a heading amongst it's children, or their children, then it's a linkable heading
    (aElement) =>
      headingAmongChildren(aElement) ||
      Array.from(aElement.children).some(headingAmongChildren)
  );

  // Scroll to the anchor in the URL when the page loads
  if (window.location.hash) {
    const hash = window.location.hash.replace("#", "");
    const heading = headings.find((x) => x.name === hash);
    if (heading) {
      heading.scrollIntoView();
    }
  }

  // Add a hash to the URL when the user scrolls to a linkable heading
  document.addEventListener("scroll", (e) => {
    headings.forEach((ha) => {
      const rect = ha.getBoundingClientRect();
      if (rect.top > 0 && rect.top < 150) {
        const location = window.location.toString().split("#")[0];
        history.replaceState(null, null, location + "#" + ha.name);
      }
    });
  });
});

function headingAmongChildren(element) {
  return (
    Array.from(element.children).find((child) =>
      ["h1", "h2", "h3"].includes(child.tagName.toLowerCase())
    ) !== undefined
  );
}
