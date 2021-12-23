
function doResponse(){
  const emailInput = document.querySelector('.emailInput');
const passwordInput = document.querySelector('.passwordInput');
const FnameInput = document.querySelector('.frNameInput');
const LnameInput = document.querySelector('.lstNameInput');

  const url = "http://127.0.0.1:8080/login";
  const data = {
    FirstName: FnameInput.value,
    LastName: LnameInput.value,
    email: emailInput.value,
    password: passwordInput.value,
  }
  const response = fetch(url, {
    method: 'POST',
    mode: 'cors', 
    cache: 'no-cache',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow', 
    referrerPolicy: 'no-referrer', 
    body: JSON.stringify(data) 
  });
}

function registration() {
const sendBut = document.querySelector('.button')
sendBut.addEventListener('click', doResponse);    
}
  window.addEventListener('load', registration);