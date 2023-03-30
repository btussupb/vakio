let popup = document.querySelector(".popup")
let overlay = document.querySelector(".overlay")
let popupOpen = document.querySelector(".popupOpen")
let closePopup = document.querySelector(".popup svg")

if (popup) {
    popupOpen.addEventListener("click", () => {
        popup.classList.toggle("active");
        overlay.classList.toggle("active");
    })

    closePopup.addEventListener("click", () => {
        popup.classList.remove("active");
        overlay.classList.remove("active");
    })
    overlay.addEventListener("click", () => {
        popup.classList.remove("active");
        overlay.classList.remove("active");
    })
    var element = document.getElementById('phone-mask');
    var maskOptions = {
        mask: '+{7}(000)000-00-00'
    };
    var mask = IMask(element, maskOptions);
    // Get the form and phone number input element
    const form = document.querySelector('form');
    const phoneInput = document.querySelector('#phone-mask');

    // Listen for form submission event
    form.addEventListener('submit', (event) => {
        // Check if the phone number is fully entered
        if (phoneInput.value.length < 16) { // 18 is the length of the mask
            // If not, prevent form submission and display an error message
            event.preventDefault();
            alert('Please enter a valid phone number');
        }
    });
}


// $('.product_slider__container').slick({
//     slidesToShow: 1,
//     slidesToScroll: 1,
//     fade: true,
//     arrows: false,
//     autoplay: true,
//     autoplaySpeed: 1000,
//     // prevArrow: '<button type="button" class="slick_arrow slick-prev"><svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M19.8779 10.1857L1.87205 10.1857M1.87205 10.1857L10.875 1.18271M1.87205 10.1857L10.875 19.1886" stroke="#fff" stroke-width="2.12202"/> </svg></button>',
//     // nextArrow: '<button type="button" class="slick_arrow slick-next"><svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M19.8779 10.1857L1.87205 10.1857M1.87205 10.1857L10.875 1.18271M1.87205 10.1857L10.875 19.1886" stroke="#fff" stroke-width="2.12202"/> </svg></button>',
//     // appendArrows: '.slider__arrow',
// });


// // alert 
let alertt = document.querySelector(".alert--fixed");
let alertClose = document.querySelectorAll(".alert--close")
let alertButton = document.querySelector(".alert-text button")
for (let item of alertClose) {
    item.addEventListener('click', function(event) {
        alertt.classList.remove("alert--active")
        alertt.classList.remove("alert--warning")
        alertt.classList.remove("alert--error")
    })
}

$(".way").waypoint({
    handler: function() {
        $(this.element).addClass("way--active");
    },
    offset: "85%",
})

function submitForm() {
    $('#form_loader').show();
}


const fileupload = document.querySelector(".fileupload")
const fileInput = document.getElementById('myFileInput');
const preview = document.getElementById('preview');

if (fileupload) {
    fileInput.addEventListener('change', () => {
      const file = fileInput.files[0];
      const reader = new FileReader();
    
      reader.addEventListener('load', () => {
        const fileType = file.type.split('/')[0];
        if (fileType === 'image') {
          preview.innerHTML = `<img src="${reader.result}" alt="preview">`;
        } else if (fileType === 'video') {
          preview.innerHTML = `<video src="${reader.result}" controls></video>`;
        }
      });
    
      reader.readAsDataURL(file);
    });
}