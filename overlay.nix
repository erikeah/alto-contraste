final:
{ system, ... }:
let
  alto-contraste = import ./package.nix { };
in
{
  inherit alto-contraste;
}
