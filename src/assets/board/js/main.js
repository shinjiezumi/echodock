let BOARD = {}
BOARD.COMMON = {}

BOARD.COMMON.EVENTS = {
  init: function () {
    this.setParameters();
    this.bindEvents();
    this.initialize();
  },
  setParameters: function () {
    this.$csrf = $('#csrf')
  },
  bindEvents: function () {
    $(document).on('click', '.deleteBoardBtn', $.proxy(this.handleDeleteBoard, null, this))
  },
  initialize: function () {
  },
  handleDeleteBoard: function (parent, e) {
    e.preventDefault()
    let form = $(e.target.closest('form'))
    form.children("input[name='csrf']").val(parent.$csrf.val())
    form.submit()
  }
};

$(function () {
  BOARD.COMMON.EVENTS.init();
});
