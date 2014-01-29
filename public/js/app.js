$(function(){
  $('#gender').keyup(function(e){
    if (e.which == 13) {
      var val = $(this).val()
      $.ajax({
        url: '/classify/' + val,
        dataType: 'json',
        success: function(data) {
          $('#gender-result').html(data.gender)
          $('#probability').html(data.probability + "%")
        }
      })
    }
  })
})