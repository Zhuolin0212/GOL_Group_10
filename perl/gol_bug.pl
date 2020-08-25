#!/usr/bin/perl

=pod
Here, X --> alive cell
      O --> dead cell

Input Pattern : 
0 0 0 0 0 
0 0 X 0 0
0 0 X 0 0
0 0 X 0 0 
0 0 0 0 0
Output Pattern :
If the generation number is odd : | If the generation number is even :
								  |
0 0 0 0 0                         | 0 0 0 0 0
0 0 0 0 0                         | 0 0 X 0 0
0 X X X 0                         | 0 0 X 0 0
0 0 0 0 0                         | 0 0 X 0 0
0 0 0 0 0                         | 0 0 0 0 0
								  |
To print for debugging use :
print <variable>;  ......To print scalar quantities like int,string
print Dumper(\@arr);  .......To print data structure like arrays,hash
=cut

use Data::Dumper;
use warnings;
no warnings 'experimental::smartmatch';

my $i = 0;
# number of generations/iterations
my $gen = 5;
chomp($gen);
# input array
my @inp_arr = ([0,0,0,0,0],[0,0,1,0,0],[0,0,1,0,0],[0,0,1,0,0],[0,0,0,0,0]);
my @checker_arr = ([0,0,0,0,0],[0,0,0,0,0],[0,1,1,1,0],[0,0,0,0,0],[0,0,0,0,0]);
my @final_arr;


## helper
sub getValue {
    my($x, $y ) = @_;
    my $s = 1;
    my $xbound = scalar @inp_arr;
    my $ybound = scalar @inp_arr;
        if ($x < 0 || $y < 0 || $x > $xbound-$s || $y > $ybound-$s) {
        return 0;
        } else {
            return($inp_arr[$x][$y]);
        }
}

# test function
sub tester {
    my @list = @_;
    $checksum = $list[0];
    if ($checksum == 99) {
        if (@final_arr ~~ @checker_arr){
            print "Tests passed\n";
        } else {
            print "Wrong patterns generated\n"
        }
    } else {
        print "\nFailed!!\n";
    }
}

## main function
sub main {

my $check = 0;
@print_arr = map { [@$_] } @inp_arr;
for (my $i = 0; $i < $gen; $i++){
    #print "Generation: $i\n";
	#print map((map($_ ? "X " : "O ", @$_), "\n"), @inp_arr);
    #print "\n\n";
    my $row = @inp_arr;
    for my $j (0 .. $inp_arr) {
        if ($j == 2){
        $check = 99;
        }
        for $k (0..$#{ $inp_arr[$i] }) { 
            my $count = 0;
        for ( [$j-1, $k-1], [$j-1, $k], [$j-1, $k+1],[ $j, $k-1],[ $j, $k+1],[ $j+1, $k-1], [ $j+1, $k], [ $j+1, $k+1]) {
            my $aa = $_->[0];
            my $bb = $_->[1];
            if (getValue($aa,$bb) == 1){
            $count++;
            }

        }
        # checker to make cell dead/alive
        if ($inp_arr[$j][$k] == 0){
                if ($count == 2 || $count == 3) {
                    $print_arr[$j][$k] = 1;
                } else {
                    $print_arr[$j][$k] = 0;
		}
        }
        # checker to make cell dead/alive
        if ($inp_arr[$j][$k] == 1) {
                if ($count == 3) {
                    $print_arr[$j][$k] = 1;
                } else {
                    $print_arr[$j][$k] = 0;
        }
        }

	}
    }

    @inp_arr = map { [@$_] } @print_arr;
    }
@final_arr = map { [@$_] } @inp_arr;
tester($check)
}
main()