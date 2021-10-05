(function () {

  const customersTable = document.querySelector(".customers__table");
  if (!customersTable) return;
  customersTable.addEventListener("click", function (evt) {
    if (evt.target.classList.contains("customers__btn-delete")) {
      const canDelete = confirm("Are you sure want to delete this customer?");
      if (!canDelete) {
        evt.preventDefault();
      }
    }
  });

})();