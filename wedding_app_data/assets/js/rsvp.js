function EventType(){
  var options = $("#eventtype option:selected");
  console.log(options.val())
  if (options.val() == "仅送祝福"){
    console.log("仅送祝福");
    $("#guest_num").hide();
    $("#RSVP").val("送祝福");
    $("#RSVP").html("送祝福");
  }
  else{
    console.log("我要出席");
    $("#guest_num").show();
    $("#RSVP").val("我要出席");
    $("#RSVP").html("我要出席");
  }
}
