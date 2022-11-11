use Sparrow6::DSL;

my $s = task-run ".", %(
  message => "Hello"
);

say "message: ", $s<Message>;
