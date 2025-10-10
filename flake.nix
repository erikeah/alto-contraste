{
  inputs = { };
  outputs =
    inputs:
    let
      forAllSystems =
        attrs:
        builtins.listToAttrs (
          builtins.map (system: {
            name = system;
            value = attrs;
          }) [ "x86_64-linux" ]
        );
    in
    {
      packages = forAllSystems {
        default = import ./package.nix { };
      };
      overlays.default = import ./overlay.nix;
    };
}
