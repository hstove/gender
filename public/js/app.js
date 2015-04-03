function titleize(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}
$(function(){
  $('#gender').keyup(function(e){
    if (e.which == 13) {
      var val = $(this).val();
      $.ajax({
        url: '/classify/' + val,
        dataType: 'json',
        success: function(data) {
          $('.result-divider').removeClass('hidden');
          var probability = new Number(data.probability).toFixed(2);
          $('#gender-result').html(data.gender);
          $('#probability').html(probability + "%");
          link = "http://gender.hankstoever.com/classify/"+val;
          $('#link a').attr('href',link).text(link);
          $('#literate').text(probability + '% of Americans named '+ titleize(val) + ' are ' + data.gender + '.');
        }
      });
    }
  });
});
