// ----------------------------------------------------------------
// Load and persist form in Local Storage.
// After call, this function immediately will try to load 
// form from LocalStorage
// ----------------------------------------------------------------
function createFormPersister(storageKey, formEl) {
  if (!formEl) throw Error(`[FormPersister]: Expected formEl to be a FormElement`);
  if (!storageKey) throw Error(`[FormPersister]: storageKey must be defined.`);

  const init = () => {
    formEl.addEventListener("submit", () => saveForm());
    formEl.addEventListener("reset", () => resetForm());
    loadForm();
  }

  const loadForm = () => {
    const formStr = window.localStorage.getItem(storageKey);
    if (formStr === null) {
      return
    }
    const keyVals = JSON.parse(formStr);
    for (let i = 0; i < formEl.elements.length; i++) {
      const element = formEl.elements[i];
      if (element.nodeName === "INPUT" && element.value === "") {
        if (keyVals[element.name] !== null) {
          element.value = keyVals[element.name];
        }
      }
    }
  };

  const saveForm = () => {
    const keyVals = {}
    for (let i = 0; i < formEl.elements.length; i++) {
      const element = formEl.elements[i];
      if (element.nodeName === "INPUT") {
        keyVals[element.name] = element.value;
      }
    }
    window.localStorage.setItem(storageKey, JSON.stringify(keyVals));
  };

  const resetForm = () => {
    window.localStorage.removeItem(storageKey);
  }

  init();

  return {
    loadForm, saveForm, resetForm
  }
}

(function () {

  // registration form
  // ------------------------------------------------------------
  function handleRegisterForm() {
    const regForm = document.querySelector(".register-form");
    if (regForm) {
      createFormPersister("register-form", regForm);
    }
  }

  // login form
  // ------------------------------------------------------------
  function handleLoginForm() {
    const logForm = document.querySelector(".login-form");
    if (logForm) {
      createFormPersister("login-form", logForm);
    }
  }

  function onDOMLoad() {
    handleRegisterForm();
    handleLoginForm();
  }

  document.addEventListener("DOMContentLoaded", onDOMLoad);
})();